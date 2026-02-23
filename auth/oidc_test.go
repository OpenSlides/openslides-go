package auth

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/OpenSlides/openslides-go/auth/authtest"
	"github.com/golang-jwt/jwt/v4"
)

// TestOIDCValidator tests OIDC token validation via mock JWKS
func TestOIDCValidator(t *testing.T) {
	// Generate test RSA key pair
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("Failed to generate RSA key: %v", err)
	}

	// Start mock JWKS server
	jwksServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwks := map[string]interface{}{
			"keys": []map[string]string{{
				"kid": "test-key-id",
				"kty": "RSA",
				"alg": "RS256",
				"n":   base64.RawURLEncoding.EncodeToString(privateKey.N.Bytes()),
				"e":   base64.RawURLEncoding.EncodeToString(big.NewInt(int64(privateKey.E)).Bytes()),
			}},
		}
		json.NewEncoder(w).Encode(jwks)
	}))
	defer jwksServer.Close()

	issuerURL := jwksServer.URL
	clientID := "test-client"

	validator := newOIDCValidator(issuerURL, issuerURL, clientID)

	t.Run("ValidToken", func(t *testing.T) {
		token := createTestOIDCToken(t, privateKey, issuerURL, clientID, "keycloak-user-123")
		claims, err := validator.ValidateToken(context.Background(), token)
		if err != nil {
			t.Fatalf("ValidateToken failed: %v", err)
		}
		if claims.KeycloakID != "keycloak-user-123" {
			t.Errorf("Expected KeycloakID 'keycloak-user-123', got '%s'", claims.KeycloakID)
		}
		if claims.Username != "testuser" {
			t.Errorf("Expected Username 'testuser', got '%s'", claims.Username)
		}
		if claims.Email != "test@example.com" {
			t.Errorf("Expected Email 'test@example.com', got '%s'", claims.Email)
		}
	})

	t.Run("ExpiredToken", func(t *testing.T) {
		token := createExpiredOIDCToken(t, privateKey, issuerURL, clientID)
		_, err := validator.ValidateToken(context.Background(), token)
		if err == nil {
			t.Error("Expected error for expired token")
		}
	})

	t.Run("WrongIssuer", func(t *testing.T) {
		token := createTestOIDCToken(t, privateKey, "https://wrong-issuer", clientID, "user")
		_, err := validator.ValidateToken(context.Background(), token)
		if err == nil {
			t.Error("Expected error for wrong issuer")
		}
	})

	t.Run("WrongAudience", func(t *testing.T) {
		token := createTestOIDCToken(t, privateKey, issuerURL, "wrong-client", "user")
		_, err := validator.ValidateToken(context.Background(), token)
		if err == nil {
			t.Error("Expected error for wrong audience")
		}
	})

	t.Run("InvalidSignature", func(t *testing.T) {
		// Create token signed with different key
		otherKey, err := rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			t.Fatalf("Failed to generate other RSA key: %v", err)
		}
		token := createTestOIDCToken(t, otherKey, issuerURL, clientID, "user")
		_, err = validator.ValidateToken(context.Background(), token)
		if err == nil {
			t.Error("Expected error for invalid signature")
		}
	})

	t.Run("MissingKid", func(t *testing.T) {
		token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
			"iss": issuerURL,
			"aud": clientID,
			"sub": "user",
			"exp": time.Now().Add(time.Hour).Unix(),
			"iat": time.Now().Unix(),
		})
		// Don't set kid in header
		signed, err := token.SignedString(privateKey)
		if err != nil {
			t.Fatalf("Failed to sign token: %v", err)
		}
		_, err = validator.ValidateToken(context.Background(), signed)
		if err == nil {
			t.Error("Expected error for missing kid")
		}
	})
}

// TestOIDCValidatorKeyCache tests that JWKS keys are properly cached
func TestOIDCValidatorKeyCache(t *testing.T) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("Failed to generate RSA key: %v", err)
	}

	requestCount := 0
	jwksServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		jwks := map[string]interface{}{
			"keys": []map[string]string{{
				"kid": "test-key-id",
				"kty": "RSA",
				"alg": "RS256",
				"n":   base64.RawURLEncoding.EncodeToString(privateKey.N.Bytes()),
				"e":   base64.RawURLEncoding.EncodeToString(big.NewInt(int64(privateKey.E)).Bytes()),
			}},
		}
		json.NewEncoder(w).Encode(jwks)
	}))
	defer jwksServer.Close()

	validator := newOIDCValidator(jwksServer.URL, jwksServer.URL, "test-client")

	// Validate multiple tokens - should only fetch JWKS once
	for i := 0; i < 5; i++ {
		token := createTestOIDCToken(t, privateKey, jwksServer.URL, "test-client", "user")
		_, err := validator.ValidateToken(context.Background(), token)
		if err != nil {
			t.Fatalf("ValidateToken failed: %v", err)
		}
	}

	if requestCount != 1 {
		t.Errorf("Expected 1 JWKS request (cached), got %d", requestCount)
	}
}

// TestOIDCTestServer tests the authtest helper
func TestOIDCTestServer(t *testing.T) {
	ots, err := authtest.NewOIDCTestServer("test-client")
	if err != nil {
		t.Fatalf("Failed to create OIDCTestServer: %v", err)
	}
	defer ots.Close()

	validator := newOIDCValidator(ots.IssuerURL, ots.IssuerURL, ots.ClientID)

	t.Run("CreateToken", func(t *testing.T) {
		token, err := ots.CreateToken("user-123")
		if err != nil {
			t.Fatalf("Failed to create token: %v", err)
		}

		claims, err := validator.ValidateToken(context.Background(), token)
		if err != nil {
			t.Fatalf("ValidateToken failed: %v", err)
		}
		if claims.KeycloakID != "user-123" {
			t.Errorf("Expected KeycloakID 'user-123', got '%s'", claims.KeycloakID)
		}
	})

	t.Run("CreateExpiredToken", func(t *testing.T) {
		token, err := ots.CreateExpiredToken("user-123")
		if err != nil {
			t.Fatalf("Failed to create expired token: %v", err)
		}

		_, err = validator.ValidateToken(context.Background(), token)
		if err == nil {
			t.Error("Expected error for expired token")
		}
	})

	t.Run("CreateTokenWithClaims", func(t *testing.T) {
		token, err := ots.CreateTokenWithClaims(jwt.MapClaims{
			"sub":   "custom-user",
			"email": "custom@example.com",
		})
		if err != nil {
			t.Fatalf("Failed to create token with claims: %v", err)
		}

		claims, err := validator.ValidateToken(context.Background(), token)
		if err != nil {
			t.Fatalf("ValidateToken failed: %v", err)
		}
		if claims.KeycloakID != "custom-user" {
			t.Errorf("Expected KeycloakID 'custom-user', got '%s'", claims.KeycloakID)
		}
		if claims.Email != "custom@example.com" {
			t.Errorf("Expected Email 'custom@example.com', got '%s'", claims.Email)
		}
	})
}

// Helper functions

func createTestOIDCToken(t *testing.T, key *rsa.PrivateKey, issuer, aud, sub string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iss":                issuer,
		"aud":                aud,
		"sub":                sub,
		"exp":                time.Now().Add(time.Hour).Unix(),
		"iat":                time.Now().Unix(),
		"preferred_username": "testuser",
		"email":              "test@example.com",
	})
	token.Header["kid"] = "test-key-id"

	signed, err := token.SignedString(key)
	if err != nil {
		t.Fatalf("Failed to sign token: %v", err)
	}
	return signed
}

func createExpiredOIDCToken(t *testing.T, key *rsa.PrivateKey, issuer, aud string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iss": issuer,
		"aud": aud,
		"sub": "user",
		"exp": time.Now().Add(-time.Hour).Unix(), // Expired
		"iat": time.Now().Add(-2 * time.Hour).Unix(),
	})
	token.Header["kid"] = "test-key-id"

	signed, err := token.SignedString(key)
	if err != nil {
		t.Fatalf("Failed to sign token: %v", err)
	}
	return signed
}
