package jwt_test

import (
	"testing"
	"time"

	"github.com/dsantaguida/idle-clicker/pkg/jwt"
)

func TestJwt(t *testing.T) {
	id := "token"

	tokenString, err := jwt.CreateToken(id)
	if err != nil {
		t.Fatal(err)
	}

	parsedId, err := jwt.ParseId(tokenString)
	if err != nil {
		t.Fatal(err)
	}
	if id != parsedId {
		t.Fatal("Incorrect id")
	}
}

func TestExpiry(t *testing.T) {
	id := "token"

	tokenString, err := jwt.CreateToken(id)
	if err != nil {
		t.Fatal(err)
	}

	err = jwt.Validate(tokenString)
	if err != nil {
		t.Fatal(err)
	}

	err = jwt.ValidateWithTime(tokenString, time.Now().Add(time.Hour * 25))
	if err == nil {
		t.Fatal("token should not have passed validation")
	}
}