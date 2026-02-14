package utils

import (
	"testing"
)

func TestGenerateHashPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		wantErr  bool
	}{
		{
			name:     "Valid password",
			password: "testpassword123",
			wantErr:  false,
		},
		{
			name:     "Empty password",
			password: "",
			wantErr:  false,
		},
		{
			name:     "Long password",
			password: "thisissomeverylongpasswordthatshouldalsowork123456789",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash, err := GenerateHashPassword(tt.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateHashPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && hash == "" {
				t.Error("GenerateHashPassword() returned empty hash")
			}
			// Verify the hash is different from the password
			if !tt.wantErr && hash == tt.password {
				t.Error("GenerateHashPassword() hash should be different from password")
			}
		})
	}
}

func TestCheckPasswordHash(t *testing.T) {
	password := "testpassword123"
	hash, err := GenerateHashPassword(password)
	if err != nil {
		t.Fatalf("Failed to generate hash: %v", err)
	}

	tests := []struct {
		name     string
		password string
		hash     string
		want     bool
	}{
		{
			name:     "Correct password",
			password: password,
			hash:     hash,
			want:     true,
		},
		{
			name:     "Incorrect password",
			password: "wrongpassword",
			hash:     hash,
			want:     false,
		},
		{
			name:     "Empty password",
			password: "",
			hash:     hash,
			want:     false,
		},
		{
			name:     "Invalid hash",
			password: password,
			hash:     "invalidhash",
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CheckPasswordHash(tt.password, tt.hash)
			if got != tt.want {
				t.Errorf("CheckPasswordHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
