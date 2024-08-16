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