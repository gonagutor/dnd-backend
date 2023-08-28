package auth_utils

import (
	utils_constants "dnd/backend/constants/utils"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type EmailTokenClaims struct {
	Type string `json:"type"`
	jwt.RegisteredClaims
}

func GenerateEmailToken(userId string) (string, error) {
	jwtSecret, _ := os.LookupEnv("JWT_SECRET")
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS512, EmailTokenClaims{
		utils_constants.EMAIL_TOKEN_TYPE,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(utils_constants.EMAIL_TOKEN_EXPIRATION * time.Hour)),
			Issuer:    utils_constants.ISSUER,
			Subject:   userId,
		},
	})
	return jwtToken.SignedString([]byte(jwtSecret))
}

func ValidateEmailToken(tokenString string) (id string, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &EmailTokenClaims{}, KeyFunc)
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*EmailTokenClaims)
	print(ok)
	if !ok || !token.Valid || claims.Subject == "" || claims.Type != utils_constants.EMAIL_TOKEN_TYPE || claims.Issuer != utils_constants.ISSUER {
		return "", errors.New("invalid token: authentication failed")
	}
	return claims.Subject, nil
}
