// Package auth implement the auth system from the openslides-auth-service:
// https://github.com/OpenSlides/openslides-auth-service
package auth

import (
	"context"
	"fmt"
	"net/http"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/OpenSlides/openslides-go/environment"
	"github.com/OpenSlides/openslides-go/oserror"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ostcar/topic"
)

var (
	envOIDCEnabled = environment.NewVariable("OIDC_ENABLED", "false", "Enable OIDC authentication.")

	// Environment variables for old system
	envAuthHost       = environment.NewVariable("AUTH_HOST", "localhost", "Host of the auth service.")
	envAuthPort       = environment.NewVariable("AUTH_PORT", "9004", "Port of the auth service.")
	envAuthProtocol   = environment.NewVariable("AUTH_PROTOCOL", "http", "Protocol of the auth service.")
	envAuthFake       = environment.NewVariable("AUTH_FAKE", "false", "Use user id 1 for every request. Ignores all other auth environment variables.")
	envAuthTokenFile  = environment.NewVariable("AUTH_TOKEN_KEY_FILE", "/run/secrets/auth_token_key", "Key to sign the JWT auth tocken.")
	envAuthCookieFile = environment.NewVariable("AUTH_COOKIE_KEY_FILE", "/run/secrets/auth_cookie_key", "Key to sign the JWT auth cookie.")

	// Environment variables for new OIDC system
	envOIDCIssuerURL         = environment.NewVariable("OIDC_ISSUER_URL", "", "Keycloak Realm URL (external) for issuer validation.")
	envOIDCInternalIssuerURL = environment.NewVariable("OIDC_INTERNAL_ISSUER_URL", "", "Keycloak Realm URL (internal) for JWKS discovery. Defaults to OIDC_ISSUER_URL.")
	envOIDCClientID          = environment.NewVariable("OIDC_CLIENT_ID", "", "Expected audience in OIDC token.")
)

// LogoutEventer tells, when a sessionID gets revoked.
//
// The method LogoutEvent has to block until there are new data. The returned
// data is a list of sessionIDs that are revoked.
type LogoutEventer interface {
	LogoutEvent(context.Context) ([]string, error)
}

type Auth interface {
	// TODO: In the new system, Authenticate does not need a ResponseWriter.
	// Remove it, when the old system can be removed.
	Authenticate(http.ResponseWriter, *http.Request) (context.Context, error)
	FromContext(context.Context) int
	AuthenticatedContext(context.Context, int) context.Context
}

// New initializes the Auth object.
func New(lookup environment.Environmenter, messageBus LogoutEventer, pool *pgxpool.Pool) (Auth, func(context.Context, func(error)), error) {
	oidcEnabled, _ := strconv.ParseBool(envOIDCEnabled.Value(lookup))
	if oidcEnabled {
		return newOIDC(lookup, messageBus, pool)
	}
	return NewOldAuthSystem(lookup, messageBus)
}

// Auth authenticates a request against the openslides-auth-service.
//
// Has to be initialized with auth.New().
type authOCID struct {
	logedoutSessions *topic.Topic[string]
	oidcValidator    *oidcValidator
	userLookup       *userLookup
}

// newOIDC initializes the Auth object with the new oidc system.
//
// Returns the initialized Auth object and a function to be called in the
// background.
//
// The pool parameter is optional and only needed for OIDC mode to lookup users
// by keycloak_id.
func newOIDC(lookup environment.Environmenter, messageBus LogoutEventer, pool *pgxpool.Pool) (*authOCID, func(context.Context, func(error)), error) {
	var oidcValidator *oidcValidator
	var userLookup *userLookup

	issuerURL := envOIDCIssuerURL.Value(lookup)
	internalIssuerURL := envOIDCInternalIssuerURL.Value(lookup)
	clientID := envOIDCClientID.Value(lookup)
	if issuerURL == "" || clientID == "" {
		return nil, nil, fmt.Errorf("OIDC enabled but OIDC_ISSUER_URL or OIDC_CLIENT_ID not set")
	}
	// Use internal URL for JWKS if provided, otherwise use issuer URL
	if internalIssuerURL == "" {
		internalIssuerURL = issuerURL
	}
	oidcValidator = newOIDCValidator(issuerURL, internalIssuerURL, clientID)

	userLookup = newUserLookup(pool)

	a := &authOCID{
		logedoutSessions: topic.New[string](),

		oidcValidator: oidcValidator,
		userLookup:    userLookup,
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
func (a *authOCID) Authenticate(_ http.ResponseWriter, r *http.Request) (context.Context, error) {
	ctx := r.Context()

	userID, sessionID, err := a.loadToken(r)
	if err != nil {
		return nil, fmt.Errorf("reading token: %w", err)
	}

	if userID == 0 {
		return a.AuthenticatedContext(ctx, 0), nil
	}

	cid, sessionIDs := a.logedoutSessions.ReceiveAll()
	if slices.Contains(sessionIDs, sessionID) {
		return nil, &authError{msg: "invalid session"}
	}

	ctx, cancelCtx := context.WithCancelCause(a.AuthenticatedContext(ctx, userID))

	go func() {
		defer cancelCtx(nil)

		var sessionIDs []string
		var err error
		for {
			cid, sessionIDs, err = a.logedoutSessions.ReceiveSince(ctx, cid)
			if err != nil {
				return
			}

			if slices.Contains(sessionIDs, sessionID) {
				cancelCtx(LogoutError{})
				return
			}
		}
	}()

	return ctx, nil
}

// AuthenticatedContext returns a new context that contains an userID.
//
// Should only used for internal URLs. All other URLs should use auth.Authenticate.
func (a *authOCID) AuthenticatedContext(ctx context.Context, userID int) context.Context {
	return context.WithValue(ctx, userIDType, userID)
}

// FromContext returnes the user id from a context returned by Authenticate().
//
// If the user is not logged in (public access) user 0 is returned.
//
// Panics, if the context was not returned from Authenticate
func (a *authOCID) FromContext(ctx context.Context) int {
	v := ctx.Value(userIDType)
	if v == nil {
		panic("call to auth.FromContext() without auth.Authenticate()")
	}

	return v.(int)
}

// listenOnLogouts listen on logout events and closes the connections.
func (a *authOCID) listenOnLogouts(ctx context.Context, logoutEventer LogoutEventer, errHandler func(error)) {
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
func (a *authOCID) pruneOldData(ctx context.Context) {
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

func (a *authOCID) loadToken(r *http.Request) (int, string, error) {
	header := r.Header.Get(authHeader)
	cookie, err := r.Cookie(cookieName)
	if err != nil && err != http.ErrNoCookie {
		return 0, "", fmt.Errorf("reading cookie: %w", err)
	}

	encodedToken := strings.TrimPrefix(header, "bearer ")
	hasBearerToken := header != encodedToken

	// No token and no auth cookie - public access
	if cookie == nil && !hasBearerToken {
		return 0, "", nil
	}

	claims, err := a.oidcValidator.ValidateToken(r.Context(), encodedToken)
	if err != nil {
		return 0, "", authError{msg: "Invalid OIDC token", wrapped: err}
	}

	// OIDC token valid - lookup user ID by keycloak_id
	userID, err := a.userLookup.GetUserIDByKeycloakID(r.Context(), claims.KeycloakID)
	if err != nil {
		return 0, "", authError{msg: fmt.Sprintf("user not found: %s", claims.KeycloakID), wrapped: err}
	}

	return userID, claims.SessionID, nil
}
