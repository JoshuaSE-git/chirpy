package auth

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "password"
	got, _ := HashPassword(password)
	if got == password {
		t.Errorf("expected different password and hash values with %q got %q", password, got)
	}
}

func TestCheckPasswordHash(t *testing.T) {
	password := "password"
	hash, _ := HashPassword(password)

	got := CheckPassword(password, hash)

	if got != nil {
		t.Errorf("expected no errors but got one")
	}
}
