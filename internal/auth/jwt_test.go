package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestJWT(t *testing.T) {
	userID := uuid.New()
	secret := "secret1"
	token, _ := MakeJWT(userID, secret, 5*time.Hour)
	tokenExpired, _ := MakeJWT(userID, secret, -5*time.Hour)

	tests := []struct {
		name      string
		id        uuid.UUID
		secret    string
		token     string
		wantError bool
	}{
		{
			name:      "valid token valid secret",
			id:        userID,
			secret:    secret,
			token:     token,
			wantError: false,
		},
		{
			name:      "valid token invalid secret",
			id:        uuid.Nil,
			secret:    "invalid secret",
			token:     token,
			wantError: true,
		},
		{
			name:      "expired token",
			id:        uuid.Nil,
			secret:    secret,
			token:     tokenExpired,
			wantError: true,
		},
		{
			name:      "expired token and invalid secret",
			id:        uuid.Nil,
			secret:    "invalid secret",
			token:     tokenExpired,
			wantError: true,
		},
		{
			name:      "invalid token",
			id:        uuid.Nil,
			secret:    secret,
			token:     "invalid.jwt.token",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := ValidateJWT(tt.token, tt.secret)
			if (err != nil) != tt.wantError {
				t.Fatalf("error: %v wantError: %v", err, tt.wantError)
			}
			if id != tt.id {
				t.Errorf("expected id: %v got: %v", tt.id, id)
			}
		})
	}
}
