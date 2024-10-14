package jwt

import (
	"context"
	"os"
	"time"

	"github.com/dsantaguida/idle-clicker/pkg/idle_errors"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/metadata"
)

const TOKEN_KEY string = "token"

type IdleClaims struct {
	Id string `json:"TOKEN"`;
	jwt.RegisteredClaims
}

func CreateToken(id string) (string, error) {
	idleClaims := &IdleClaims{
		id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, idleClaims)

	tokenString, err := token.SignedString([]byte(os.Getenv("SIGNING_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseId(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &IdleClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNING_KEY")), nil
	},
		jwt.WithExpirationRequired(),
		jwt.WithIssuedAt())
	if err != nil {
		return "", err
	} 
	
	claims, ok := token.Claims.(*IdleClaims)
	if !ok {
		return "", idle_errors.ErrUnknownClaimsType
	}

	return claims.Id,  nil
}

func ValidateWithTime(tokenString string, currentTime time.Time) error {
	token, err := jwt.ParseWithClaims(tokenString, &IdleClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNING_KEY")), nil
	})
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(*IdleClaims)
	if !ok {
		return idle_errors.ErrUnknownClaimsType
	}

	expTime, err := claims.GetExpirationTime()
	if err != nil {
		return err
	}

	if currentTime.After(expTime.Time) {
		return idle_errors.ErrExpiredToken
	}

	return nil
}

func Validate(tokenString string) error {
	return ValidateWithTime(tokenString, time.Now())
}

func GetTokenFromContext(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", idle_errors.ErrNoMetadata
	}

	token := md.Get(TOKEN_KEY)[0]
	if len(token) == 0 {
		return "", idle_errors.ErrTokenNotInHeader
	}

	return token, nil
}