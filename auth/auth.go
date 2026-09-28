// Package auth implement the auth system from the openslides-auth-service:
// https://github.com/OpenSlides/openslides-auth-service
package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/OpenSlides/openslides-go/environment"
	"github.com/golang-jwt/jwt/v4"
)

var (
	envExternalHost = environment.NewVariable("IDP_EXTERNAL_HOST", "localhost:8800", "External host address")
	logoutMU        sync.RWMutex
)

// pruneTime defines how long a topic id will be valid. This should be higher
// than the max lifetime of a token.
const pruneTime = 15 * time.Minute

const (
	authHeader = "Authorization"
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

	ll, err := NewLogoutListener(lookup)
	if err != nil {
		return nil, nil, err
	}

	a := &Auth{
		externalHost:   externalHost,
		logoutListener: ll,
	}

	background := func(ctx context.Context, errorHandler func(error)) {
		go a.listenOnLogouts(ctx, errorHandler)
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

	// Blocklist
	//cid, sessionIDs := a.logedoutSessions.ReceiveAll()
	//if slices.Contains(sessionIDs, p.SessionID) {
	//	return nil, &authError{"invalid session", nil}
	//}

	// Get OS User Id linked to IDP ID
	ctx, _ = context.WithCancel(a.AuthenticatedContext(ctx, p.OSUserID))

	/*
		go func() {
			defer cancelCtx()

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
			}
		}()
	*/

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

	if _, err := jwt.ParseWithClaims(encodedToken, payload, func(token *jwt.Token) (interface{}, error) {
		return nil, errors.New("token validation is handled by proxy")
	}); err != nil {
		var invalid *jwt.ValidationError
		if errors.As(err, &invalid) {
			return authError{msg: "Invalid auth token", wrapped: invalid}
		}
	}

	return nil
}

type payloadIDP struct {
	jwt.RegisteredClaims
	IDPID     string `json:"sub"`
	Issuer    string `json:"iss"`
	SessionID string `json:"sid"` // IDP session ID
	OSUserID  int    `json:"os_id"`
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
func (a *Auth) listenOnLogouts(ctx context.Context, errHandler func(error)) {
	if errHandler == nil {
		errHandler = func(error) {}
	}

	for {
		select {
		// case <-ctx.Err():
		case <-ctx.Done():
			return
		}

		logoutMU.Lock()
		// Datastore Fetch
		logoutMU.Unlock()
		time.Sleep(time.Second)
	}
}

type authString string

const (
	userIDType authString = "user_id"
)
