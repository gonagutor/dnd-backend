package auth_utils

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func KeyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.New("unexpected signing method in auth token")
	}
	jwtSecret, _ := os.LookupEnv("JWT_SECRET")
	return []byte(jwtSecret), nil
}
