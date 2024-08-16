package test

import (
	"testing"

	"github.com/dsantaguida/idle-clicker/services/authentication/internal/password"
)

func TestPasswordCompliance(t *testing.T) {
	pw := "p"
	err := password.CheckPasswordCompliance(pw)
	if err == nil {
		t.Fatal("Password should not have passed compliance check.")
	}

	pw = "password123"
	err = password.CheckPasswordCompliance(pw)
	if err != nil {
		t.Fatal("Password should have passed compliance check.")
	}
}

func TestPasswordHashing(t *testing.T) {
	pw := "password123"

	hash, err := password.HashPassword(pw)
	if err != nil {
		t.Fatal("Failed to hash password: ", err)
	}
	if len(hash) == 0 {
		t.Fatal("Failed to has password.")
	}
	
	result := password.VerifyPassword(pw, hash)
	if !result {
		t.Fatal("Failed to verify password.")
	}
}