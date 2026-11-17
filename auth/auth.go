// Package auth implement the auth system from the openslides-auth-service:
// https://github.com/OpenSlides/openslides-auth-service
package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/OpenSlides/openslides-go/environment"
	"github.com/golang-jwt/jwt/v4"
)

var (
	envExternalHost = environment.NewVariable("IDP_EXTERNAL_HOST", "localhost:8800", "External host address")
)

type authString string

// pruneTime defines how long a topic id will be valid. This should be higher
// than the max lifetime of a token.
const pruneTime = 15 * time.Minute

const (
	authHeader                 = "Authorization"
	sessionIDHeader            = "X-OIDC-Session"
	userIDType      authString = "user_id"
)

// Auth authenticates a request against the idp service.
//
// Has to be initialized with auth.New().
type Auth struct {
	externalHost   string
	logoutListener *LogoutListener
}

// New initializes the Auth object.
//
// Returns the initialized Auth object and a function to be called in the
// background.
func New(lookup environment.Environmenter) (*Auth, func(context.Context, func(error)), error) {
	externalHost := envExternalHost.Value(lookup)

	ll, background, err := NewLogoutListener(lookup)
	if err != nil {
		return nil, nil, err
	}

	// Close pool if an error occurs prior to returning background func
	defer func() {
		if r := recover(); r != nil {
			ll.pool.Close()
		}
	}()

	a := &Auth{
		externalHost:   externalHost,
		logoutListener: ll,
	}

	return a, background, nil
}

// Authenticate uses the headers from the given request to get the user id. The
// returned context will be cancled, if the session is revoked.
func (a *Auth) Authenticate(w http.ResponseWriter, r *http.Request) (context.Context, error) {
	ctx := r.Context()
	p := new(payloadIDPAccessToken)
	if err := a.parseAccessToken(r, p); err != nil {
		return nil, fmt.Errorf("reading token: %w", err)
	}

	sid := r.Header.Get(sessionIDHeader)

	if p.IDPID == "" {
		return a.AuthenticatedContext(ctx, 0), nil
	}

	// Blocklist
	if sid == "" || a.logoutListener.IsBlocked(sid) {
		return nil, &authError{"invalid session", nil}
	}

	// Get OS User Id linked to IDP ID
	ctx, cancelCtx := context.WithCancel(a.AuthenticatedContext(ctx, p.OSUserID))

	// Periodically check if the session has been blocked. If session is in blocklist, cancel the context
	go func() {
		defer func() {
			cancelCtx()
		}()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				a.logoutListener.notifyMU.Lock()
				if a.logoutListener.IsBlocked(sid) {
					a.logoutListener.notifyMU.Unlock()
					return
				}

				a.logoutListener.NotifyCond().Wait()
				a.logoutListener.notifyMU.Unlock()
			}
		}
	}()

	return ctx, nil
}

// loadToken loads and validates the token. If the token is expired, it tries
// to renew it and write the new token in the responsewriter.
func (a *Auth) parseAccessToken(r *http.Request, payload jwt.Claims) error {
	header := r.Header.Get(authHeader)
	encodedToken := strings.TrimPrefix(header, "Bearer: ")

	if header == encodedToken {
		// No token. Handle the request as public access requst.
		return nil
	}
	if _, _, err := new(jwt.Parser).ParseUnverified(encodedToken, payload); err != nil {
		var invalid *jwt.ValidationError
		if errors.As(err, &invalid) {
			return authError{msg: fmt.Sprintf("Couldn't parse JWT access token %v", err), wrapped: invalid}
		}
	}

	return nil
}

type payloadIDPAccessToken struct {
	jwt.RegisteredClaims
	IDPID    string `json:"sub"`
	Issuer   string `json:"iss"`
	OSUserID int    `json:"os_id"`
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
