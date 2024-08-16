package jwt_test

import (
	"testing"

	"github.com/dsantaguida/idle-clicker/pkg/jwt"
)

func TestJwt(t *testing.T) {
	id := "token"

	tokenString, err := jwt.CreateToken(id)
	if err != nil {
		t.Fatal(err)
	}

	claims, err := jwt.Verify(tokenString)
	if err != nil {
		t.Fatal(err)
	}
	if claims.Id != id {
		t.Fatal("Incorrect id")
	}
}