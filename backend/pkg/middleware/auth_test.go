package middleware

import (
	"testing"
)

func TestAuthConfig(t *testing.T) {
	config := DefaultAuthConfig
	
	if config.TokenLookup != "header:Authorization" {
		t.Errorf("Expected default TokenLookup to be 'header:Authorization', got %s", config.TokenLookup)
	}
}

func TestIsWhiteListed(t *testing.T) {
	tests := []struct {
		name      string
		path      string
		whiteList []string
		want      bool
	}{
		{
			name:      "Exact match",
			path:      "/api/login",
			whiteList: []string{"/api/login", "/api/register"},
			want:      true,
		},
		{
			name:      "No match",
			path:      "/api/protected",
			whiteList: []string{"/api/login", "/api/register"},
			want:      false,
		},
		{
			name:      "Prefix match with wildcard",
			path:      "/api/public/resource",
			whiteList: []string{"/api/public/*"},
			want:      true,
		},
		{
			name:      "Prefix match - subdirectory",
			path:      "/api/public/deep/resource",
			whiteList: []string{"/api/public/*"},
			want:      true,
		},
		{
			name:      "No match - similar path",
			path:      "/api/private/resource",
			whiteList: []string{"/api/public/*"},
			want:      false,
		},
		{
			name:      "Empty whitelist",
			path:      "/api/test",
			whiteList: []string{},
			want:      false,
		},
		{
			name:      "Multiple wildcards",
			path:      "/api/public/test",
			whiteList: []string{"/api/admin/*", "/api/public/*", "/api/guest/*"},
			want:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isWhiteListed(tt.path, tt.whiteList)
			if got != tt.want {
				t.Errorf("isWhiteListed(%s, %v) = %v, want %v", tt.path, tt.whiteList, got, tt.want)
			}
		})
	}
}

func TestNewAuth(t *testing.T) {
	middleware := NewAuth()
	if middleware == nil {
		t.Error("NewAuth() should return a middleware function")
	}
}

func TestNewAuthWithConfig(t *testing.T) {
	tests := []struct {
		name   string
		config AuthConfig
	}{
		{
			name:   "Default config",
			config: DefaultAuthConfig,
		},
		{
			name: "Custom config with whitelist",
			config: AuthConfig{
				TokenLookup: "header:Authorization",
				WhiteList:   []string{"/api/public/*"},
			},
		},
		{
			name: "Empty token lookup - should use default",
			config: AuthConfig{
				TokenLookup: "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			middleware := NewAuthWithConfig(tt.config)
			if middleware == nil {
				t.Error("NewAuthWithConfig() should return a middleware function")
			}
		})
	}
}
