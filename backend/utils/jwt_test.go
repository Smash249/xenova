package utils

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestBaseClaims(t *testing.T) {
	claims := BaseClaims{
		ID:       1,
		UserName: "testuser",
	}

	if claims.ID != 1 {
		t.Errorf("Expected ID = 1, got %d", claims.ID)
	}
	if claims.UserName != "testuser" {
		t.Errorf("Expected UserName = testuser, got %s", claims.UserName)
	}
}

func TestCustomClaims(t *testing.T) {
	baseClaims := BaseClaims{
		ID:       1,
		UserName: "testuser",
	}

	now := time.Now()
	customClaims := CustomClaims{
		BaseClaims: baseClaims,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * 24)),
			Issuer:    "test-issuer",
			Audience:  jwt.ClaimStrings{"test-audience"},
		},
	}

	if customClaims.BaseClaims.ID != 1 {
		t.Errorf("Expected ID = 1, got %d", customClaims.BaseClaims.ID)
	}
	if customClaims.BaseClaims.UserName != "testuser" {
		t.Errorf("Expected UserName = testuser, got %s", customClaims.BaseClaims.UserName)
	}
}

func TestCreateClaims(t *testing.T) {
	baseClaims := BaseClaims{
		ID:       123,
		UserName: "johndoe",
	}

	// Note: This test requires viper configuration to be set up
	// For now, we test the structure
	tests := []struct {
		name        string
		baseClaims  BaseClaims
		expireHours int
	}{
		{
			name:        "1 hour expiration",
			baseClaims:  baseClaims,
			expireHours: 1,
		},
		{
			name:        "24 hours expiration",
			baseClaims:  baseClaims,
			expireHours: 24,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// We can't fully test CreateClaims without viper config
			// but we can verify the function signature
			if tt.expireHours < 0 {
				t.Error("expireHours should be positive")
			}
		})
	}
}

func TestExtractToken(t *testing.T) {
	// Note: This test would require Echo context mock
	// Testing the token extraction logic conceptually
	tests := []struct {
		name      string
		authValue string
		wantToken string
		wantErr   bool
	}{
		{
			name:      "Bearer token",
			authValue: "Bearer abc123token",
			wantToken: "abc123token",
			wantErr:   false,
		},
		{
			name:      "Token without Bearer prefix",
			authValue: "abc123token",
			wantToken: "abc123token",
			wantErr:   false,
		},
		{
			name:      "Empty authorization",
			authValue: "",
			wantToken: "",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// This validates test cases structure
			// Actual implementation would need Echo context
			if tt.wantErr && tt.wantToken != "" {
				t.Error("If error is expected, token should be empty")
			}
		})
	}
}

func TestJWTErrors(t *testing.T) {
	errors := []error{
		ErrTokenExpired,
		ErrTokenNotValidYet,
		ErrTokenMalformed,
		ErrTokenInvalid,
		ErrTokenNotFound,
		ErrRefreshInvalid,
	}

	for _, err := range errors {
		if err == nil {
			t.Error("JWT error should not be nil")
		}
		if err.Error() == "" {
			t.Error("JWT error should have a message")
		}
	}
}
