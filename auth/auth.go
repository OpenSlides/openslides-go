// Package auth implement the auth system from the openslides-auth-service:
// https://github.com/OpenSlides/openslides-auth-service
package auth

import (
	"context"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/OpenSlides/openslides-go/environment"
	"github.com/OpenSlides/openslides-go/oserror"
	"github.com/golang-jwt/jwt/v4"
	"github.com/ostcar/topic"
)

var (
	envExternalHost      = environment.NewVariable("IDP_EXTERNAL_HOST", "localhost:8800", "External host address")
	envIssuerURL         = environment.NewVariable("IDP_URL_EXTERNAL", "https://localhost:8800", "URL of idp server")
	envIssuerURLInternal = environment.NewVariable("IDP_URL_INTERNAL", "http://zitadel-api:8080", "Internal URL of idp server")
)

// pruneTime defines how long a topic id will be valid. This should be higher
// than the max lifetime of a token.
const pruneTime = 15 * time.Minute

const (
	authHeader = "Authorization"
)

// LogoutEventer tells, when a sessionID gets revoked.
//
// The method LogoutEvent has to block until there are new data. The returned
// data is a list of sessionIDs that are revoked.
type LogoutEventer interface {
	LogoutEvent(context.Context) ([]string, error)
}

// Auth authenticates a request against the openslides-auth-service.
//
// Has to be initialized with auth.New().
type Auth struct {
	logedoutSessions *topic.Topic[string]

	externalHost      string
	issuerURL         string
	issuerURLInternal string

	keys      map[string]*rsa.PublicKey
	expiresAt time.Time
}

// New initializes the Auth object.
//
// Returns the initialized Auth objectand a function to be called in the
// background.
func New(lookup environment.Environmenter, messageBus LogoutEventer) (*Auth, func(context.Context, func(error)), error) {
	externalHost := envExternalHost.Value(lookup)
	issuerURL := envIssuerURL.Value(lookup)

	issuerURLInternal := envIssuerURLInternal.Value(lookup)
	if issuerURLInternal == "" {
		issuerURLInternal = envIssuerURL.Value(lookup)
	}

	a := &Auth{
		logedoutSessions:  topic.New[string](),
		externalHost:      externalHost,
		issuerURL:         issuerURL,
		issuerURLInternal: issuerURLInternal,
		keys:              make(map[string]*rsa.PublicKey),
	}

	// Make sure the topic is not empty
	a.logedoutSessions.Publish("")

	background := func(ctx context.Context, errorHandler func(error)) {
		go a.listenOnLogouts(ctx, messageBus, errorHandler)
		go a.pruneOldData(ctx)
	}

	return a, background, nil
}

// Authenticate uses the headers from the given request to get the user id. The
// returned context will be cancled, if the session is revoked.
func (a *Auth) Authenticate(w http.ResponseWriter, r *http.Request) (context.Context, error) {
	ctx := r.Context()

	p := new(payloadIDP)
	if err := a.loadTokenIDP(w, r, p); err != nil {
		return nil, fmt.Errorf("reading token: %w", err)
	}

	if p.IDPID == "" {
		return a.AuthenticatedContext(ctx, 0), nil
	}

	// TODO: Blocklist
	//cid, sessionIDs := a.logedoutSessions.ReceiveAll()
	//if slices.Contains(sessionIDs, p.SessionID) {
	//	return nil, &authError{"invalid session", nil}
	//}

	// Get OS User Id linked to IDP ID
	ctx, cancelCtx := context.WithCancel(a.AuthenticatedContext(ctx, p.OSUserID))

	go func() {
		defer cancelCtx()

		/*
			var sessionIDs []string
			var err error
			for {
				cid, sessionIDs, err = a.logedoutSessions.ReceiveSince(ctx, cid)
				if err != nil {
					return
				}

				if slices.Contains(sessionIDs, p.SessionID) {
					return
				}
			}*/
	}()

	return ctx, nil
}

// loadToken loads and validates the token. If the token is expired, it tries
// to renew it and write the new token in the responsewriter.
func (a *Auth) loadTokenIDP(w http.ResponseWriter, r *http.Request, payload jwt.Claims) error {
	header := r.Header.Get(authHeader)
	encodedToken := strings.TrimPrefix(header, "Bearer: ")

	if header == encodedToken {
		// No token. Handle the request as public access requst.
		return nil
	}

	if _, err := a.validateToken(r.Context(), encodedToken); err != nil {
		// OIDC validation failed - return error (no cookie means no legacy fallback)
		return authError{msg: "Invalid OIDC token", wrapped: err}
	}

	if _, err := jwt.ParseWithClaims(encodedToken, payload, func(token *jwt.Token) (interface{}, error) {
		rawKid, ok := token.Header["kid"]
		if !ok {
			return nil, fmt.Errorf("missing kid in token header")
		}

		kid, ok := rawKid.(string)
		if !ok {

			return nil, fmt.Errorf("kid in token header is not a string")
		}

		return a.getKey(r.Context(), kid)
	}); err != nil {
		var invalid *jwt.ValidationError
		if errors.As(err, &invalid) {
			return a.handleInvalidToken(r.Context(), invalid, w, encodedToken)
		}
	}

	return nil
}

func (a *Auth) validateToken(ctx context.Context, tokenString string) (*payloadIDP, error) {
	// 1. Parse token without validation to get kid
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, &payloadIDP{})
	if err != nil {
		return nil, fmt.Errorf("parsing token: %w", err)
	}

	rawKid, ok := token.Header["kid"]
	if !ok {
		return nil, fmt.Errorf("missing kid in token header")
	}

	kid, ok := rawKid.(string)
	if !ok {
		return nil, fmt.Errorf("kid in token header is not a string")
	}

	// 2. Get public key from JWKS (cached)
	key, err := a.getKey(ctx, kid)
	if err != nil {
		return nil, fmt.Errorf("getting key: %w", err)
	}

	// 3. Validate token with public key
	claims := &payloadIDP{}
	token, err = jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {

		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		return nil, fmt.Errorf("validating token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// 4. Validate issuer
	if claims.Issuer != a.issuerURL {
		return nil, fmt.Errorf("invalid issuer: got %s, want %s", claims.Issuer, a.issuerURL)
	}

	return claims, nil
}

type payloadIDP struct {
	jwt.RegisteredClaims
	IDPID  string `json:"sub"`
	Issuer string `json:"iss"`
	// SessionID string `json:"sid"` // IDP session ID
	OSUserID int `json:"os_id"`
}

// getKey returns the RSA public key for the given kid, fetching from JWKS if needed
func (a *Auth) getKey(ctx context.Context, kid string) (*rsa.PublicKey, error) {
	if key, ok := a.keys[kid]; ok && time.Now().Before(a.expiresAt) {
		return key, nil
	}

	// Fetch JWKS
	return a.fetchJWKS(ctx, kid)
}

// fetchJWKS fetches the JWKS from the issuer and caches the keys
func (a *Auth) fetchJWKS(ctx context.Context, kid string) (*rsa.PublicKey, error) {
	// Double-check after acquiring write lock
	if key, ok := a.keys[kid]; ok && time.Now().Before(a.expiresAt) {
		return key, nil
	}

	req, err := http.NewRequestWithContext(ctx, "GET", a.issuerURLInternal+"/oauth/v2/keys", nil)
	if err != nil {
		return nil, fmt.Errorf("creating JWKS request: %w", err)
	}

	req.Host = a.externalHost

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fetching JWKS: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("JWKS request failed: %s", resp.Status)
	}

	var jwks struct {
		Keys []struct {
			Kid string `json:"kid"`
			Kty string `json:"kty"`
			Alg string `json:"alg"`
			N   string `json:"n"`
			E   string `json:"e"`
		} `json:"keys"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&jwks); err != nil {
		return nil, fmt.Errorf("decoding JWKS: %w", err)
	}

	// Parse and cache all keys
	a.keys = make(map[string]*rsa.PublicKey)
	for _, k := range jwks.Keys {
		if k.Kty != "RSA" {
			continue
		}
		key, err := parseRSAPublicKey(k.N, k.E)
		if err != nil {
			continue
		}
		a.keys[k.Kid] = key
	}

	// Cache for 1 hour
	a.expiresAt = time.Now().Add(time.Hour)

	key, ok := a.keys[kid]
	if !ok {
		return nil, fmt.Errorf("key %s not found in JWKS", kid)
	}
	return key, nil
}

// parseRSAPublicKey parses RSA public key from JWKS n and e values
func parseRSAPublicKey(nStr, eStr string) (*rsa.PublicKey, error) {
	nBytes, err := base64.RawURLEncoding.DecodeString(nStr)
	if err != nil {
		return nil, fmt.Errorf("decoding n: %w", err)
	}
	eBytes, err := base64.RawURLEncoding.DecodeString(eStr)
	if err != nil {
		return nil, fmt.Errorf("decoding e: %w", err)
	}

	n := new(big.Int).SetBytes(nBytes)
	e := int(new(big.Int).SetBytes(eBytes).Int64())

	return &rsa.PublicKey{N: n, E: e}, nil
}

// AuthenticatedContext returns a new context that contains a userID.
//
// Should only used for internal URLs. All other URLs should use auth.Authenticate.
func (a *Auth) AuthenticatedContext(ctx context.Context, userID int) context.Context {
	return context.WithValue(ctx, userIDType, userID)
}

// FromContext returnes the user id from a context returned by Authenticate().
//
// If the user is not logged in (public access) user 0 is returned.
//
// Panics, if the context was not returned from Authenticate
func (a *Auth) FromContext(ctx context.Context) int {
	v := ctx.Value(userIDType)
	if v == nil {
		panic("call to auth.FromContext() without auth.Authenticate()")
	}

	return v.(int)
}

// listenOnLogouts listen on logout events and closes the connections.
func (a *Auth) listenOnLogouts(ctx context.Context, logoutEventer LogoutEventer, errHandler func(error)) {
	if errHandler == nil {
		errHandler = func(error) {}
	}

	for {
		data, err := logoutEventer.LogoutEvent(ctx)
		if err != nil {
			if oserror.ContextDone(err) {
				return
			}

			errHandler(fmt.Errorf("receiving logout event: %w", err))
			time.Sleep(time.Second)
			continue
		}

		a.logedoutSessions.Publish(data...)
	}
}

// pruneOldData removes old logout events.
func (a *Auth) pruneOldData(ctx context.Context) {
	tick := time.NewTicker(5 * time.Minute)
	defer tick.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
			a.logedoutSessions.Prune(time.Now().Add(-pruneTime))
		}
	}
}

func (a *Auth) handleInvalidToken(ctx context.Context, invalid *jwt.ValidationError, w http.ResponseWriter, encodedToken string) error {
	if !tokenExpired(invalid.Errors) {
		return authError{"Invalid auth token", invalid}
	}

	return authError{"Token expired", invalid}
}

func tokenExpired(errNo uint32) bool {
	return errNo&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0
}

type authString string

const (
	userIDType authString = "user_id"
)
