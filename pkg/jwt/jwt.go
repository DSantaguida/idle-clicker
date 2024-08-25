package jwt

import (
	"os"
	"time"

	"github.com/dsantaguida/idle-clicker/pkg/idle_errors"
	"github.com/golang-jwt/jwt/v5"
)

type IdleClaims struct {
	Id string `json:"TOKEN"`;
	Date string `json:"DATE"`;
	jwt.RegisteredClaims
}

func CreateToken(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"TOKEN": id,
		"DATE": time.Now(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SIGNING_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func Verify(tokenString string) (*IdleClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &IdleClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNING_KEY")), nil
	})
	if err != nil {
		return nil, err
	} else if claims, ok := token.Claims.(*IdleClaims); ok {
		return claims,  nil
	} else {
		return nil, idle_errors.ErrUnknownClaimsType
	}	
}
