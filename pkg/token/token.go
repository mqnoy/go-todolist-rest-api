package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrTokenNotValid = fmt.Errorf("token method is not valid")
)

type CustomClaimOptions struct {
	ExpiredTime *jwt.NumericDate
	SubjectId   string
}

func Generate(mapClaims jwt.MapClaims, key []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)

	tokenString, err := token.SignedString(key)

	return tokenString, err
}

func Verify(mapClaims jwt.MapClaims, key []byte, tokenString string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, mapClaims, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, ErrTokenNotValid
		}

		return key, nil
	})
}

func GenerateMapClaims(options CustomClaimOptions) jwt.MapClaims {
	return jwt.MapClaims{
		"exp": options.ExpiredTime,
		"iat": jwt.NewNumericDate(time.Now()),
		"sub": options.SubjectId,
	}
}
