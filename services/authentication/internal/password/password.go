package password

import (
	"github.com/dsantaguida/idle-clicker/pkg/idle_errors"
	"golang.org/x/crypto/bcrypt"
)

const lengthRequirement = 8

func CheckPasswordCompliance(password string) error {
	if len(password) < lengthRequirement {
		return idle_errors.ErrPasswordNotCompliant
	}
	return nil
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func VerifyPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}