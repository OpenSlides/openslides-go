package authtest

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// OIDCTestServer provides a mock Keycloak JWKS endpoint for testing
type OIDCTestServer struct {
	Server     *httptest.Server
	PrivateKey *rsa.PrivateKey
	IssuerURL  string
	ClientID   string
}

// NewOIDCTestServer creates a test server with mock JWKS endpoint
func NewOIDCTestServer(clientID string) (*OIDCTestServer, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	ots := &OIDCTestServer{
		PrivateKey: privateKey,
		ClientID:   clientID,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/protocol/openid-connect/certs", ots.handleJWKS)

	ots.Server = httptest.NewServer(mux)
	ots.IssuerURL = ots.Server.URL

	return ots, nil
}

func (ots *OIDCTestServer) handleJWKS(w http.ResponseWriter, r *http.Request) {
	jwks := map[string]interface{}{
		"keys": []map[string]string{{
			"kid": "test-key-id",
			"kty": "RSA",
			"alg": "RS256",
			"n":   base64.RawURLEncoding.EncodeToString(ots.PrivateKey.N.Bytes()),
			"e":   base64.RawURLEncoding.EncodeToString(big.NewInt(int64(ots.PrivateKey.E)).Bytes()),
		}},
	}
	json.NewEncoder(w).Encode(jwks)
}

// CreateToken creates a valid OIDC token for testing
func (ots *OIDCTestServer) CreateToken(keycloakID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iss":                ots.IssuerURL,
		"aud":                ots.ClientID,
		"sub":                keycloakID,
		"exp":                time.Now().Add(time.Hour).Unix(),
		"iat":                time.Now().Unix(),
		"preferred_username": "testuser",
		"email":              "test@example.com",
	})
	token.Header["kid"] = "test-key-id"
	return token.SignedString(ots.PrivateKey)
}

// CreateTokenWithClaims creates an OIDC token with custom claims for testing
func (ots *OIDCTestServer) CreateTokenWithClaims(claims jwt.MapClaims) (string, error) {
	// Set defaults if not provided
	if _, ok := claims["iss"]; !ok {
		claims["iss"] = ots.IssuerURL
	}
	if _, ok := claims["aud"]; !ok {
		claims["aud"] = ots.ClientID
	}
	if _, ok := claims["exp"]; !ok {
		claims["exp"] = time.Now().Add(time.Hour).Unix()
	}
	if _, ok := claims["iat"]; !ok {
		claims["iat"] = time.Now().Unix()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token.Header["kid"] = "test-key-id"
	return token.SignedString(ots.PrivateKey)
}

// CreateExpiredToken creates an expired OIDC token for testing
func (ots *OIDCTestServer) CreateExpiredToken(keycloakID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iss":                ots.IssuerURL,
		"aud":                ots.ClientID,
		"sub":                keycloakID,
		"exp":                time.Now().Add(-time.Hour).Unix(), // Expired
		"iat":                time.Now().Add(-2 * time.Hour).Unix(),
		"preferred_username": "testuser",
		"email":              "test@example.com",
	})
	token.Header["kid"] = "test-key-id"
	return token.SignedString(ots.PrivateKey)
}

// Close shuts down the test server
func (ots *OIDCTestServer) Close() {
	ots.Server.Close()
}
