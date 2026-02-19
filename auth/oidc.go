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
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// OIDCValidator validates Keycloak/OIDC tokens via JWKS
type OIDCValidator struct {
	issuerURL string
	clientID  string
	jwksURL   string

	mu        sync.RWMutex
	keys      map[string]*rsa.PublicKey
	expiresAt time.Time
}

// NewOIDCValidator creates a new OIDC token validator
// issuerURL is used for validating the token's issuer claim (external URL)
// internalIssuerURL is used for fetching JWKS (can be internal/docker network URL)
func NewOIDCValidator(issuerURL, internalIssuerURL, clientID string) *OIDCValidator {
	return &OIDCValidator{
		issuerURL: issuerURL,
		clientID:  clientID,
		jwksURL:   internalIssuerURL + "/protocol/openid-connect/certs",
		keys:      make(map[string]*rsa.PublicKey),
	}
}

// OIDCClaims represents Keycloak token claims
type OIDCClaims struct {
	jwt.RegisteredClaims
	KeycloakID      string `json:"sub"`
	SessionID       string `json:"sid"` // Keycloak session ID
	Email           string `json:"email"`
	Username        string `json:"preferred_username"`
	AuthorizedParty string `json:"azp"` // Keycloak puts client_id here instead of aud
}

// ValidateToken validates a Keycloak JWT token and returns claims
func (v *OIDCValidator) ValidateToken(ctx context.Context, tokenString string) (*OIDCClaims, error) {
	// 1. Parse token without validation to get kid
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, &OIDCClaims{})
	if err != nil {
		return nil, fmt.Errorf("parsing token: %w", err)
	}

	kid, ok := token.Header["kid"].(string)
	if !ok {
		return nil, errors.New("missing kid in token header")
	}

	// 2. Get public key from JWKS (cached)
	key, err := v.getKey(ctx, kid)
	if err != nil {
		return nil, fmt.Errorf("getting key: %w", err)
	}

	// 3. Validate token with public key
	claims := &OIDCClaims{}
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
		return nil, errors.New("invalid token")
	}

	// 4. Validate issuer and audience
	if claims.Issuer != v.issuerURL {
		return nil, fmt.Errorf("invalid issuer: got %s, want %s", claims.Issuer, v.issuerURL)
	}

	if !v.verifyAudience(claims) {
		return nil, errors.New("invalid audience")
	}

	return claims, nil
}

// verifyAudience checks if the clientID is in the token's audience or authorized party (azp)
// Keycloak typically puts the client ID in azp rather than aud for access tokens
func (v *OIDCValidator) verifyAudience(claims *OIDCClaims) bool {
	// Check azp claim (Keycloak's authorized party)
	if claims.AuthorizedParty == v.clientID {
		return true
	}
	// Check aud claim as fallback
	for _, aud := range claims.Audience {
		if aud == v.clientID {
			return true
		}
	}
	return false
}

// getKey returns the RSA public key for the given kid, fetching from JWKS if needed
func (v *OIDCValidator) getKey(ctx context.Context, kid string) (*rsa.PublicKey, error) {
	v.mu.RLock()
	if key, ok := v.keys[kid]; ok && time.Now().Before(v.expiresAt) {
		v.mu.RUnlock()
		return key, nil
	}
	v.mu.RUnlock()

	// Fetch JWKS
	return v.fetchJWKS(ctx, kid)
}

// fetchJWKS fetches the JWKS from the issuer and caches the keys
func (v *OIDCValidator) fetchJWKS(ctx context.Context, kid string) (*rsa.PublicKey, error) {
	v.mu.Lock()
	defer v.mu.Unlock()

	// Double-check after acquiring write lock
	if key, ok := v.keys[kid]; ok && time.Now().Before(v.expiresAt) {
		return key, nil
	}

	req, err := http.NewRequestWithContext(ctx, "GET", v.jwksURL, nil)
	if err != nil {
		return nil, fmt.Errorf("creating JWKS request: %w", err)
	}

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
	v.keys = make(map[string]*rsa.PublicKey)
	for _, k := range jwks.Keys {
		if k.Kty != "RSA" {
			continue
		}
		key, err := parseRSAPublicKey(k.N, k.E)
		if err != nil {
			continue
		}
		v.keys[k.Kid] = key
	}

	// Cache for 1 hour
	v.expiresAt = time.Now().Add(time.Hour)

	key, ok := v.keys[kid]
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
