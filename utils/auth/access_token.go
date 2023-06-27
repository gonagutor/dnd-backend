package auth_utils

import (
	"errors"
	"os"
	utils_constants "revosearch/backend/constants/utils"
	"revosearch/backend/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AccessTokenClaims struct {
	Type string `json:"type"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(user *models.User) (string, error) {
	jwtSecret, _ := os.LookupEnv("JWT_SECRET")
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS512, AccessTokenClaims{
		utils_constants.ACCESS_TOKEN_TYPE,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(utils_constants.ACCESS_TOKEN_EXPIRATION * time.Minute)),
			Issuer:    utils_constants.ISSUER,
			Subject:   user.ID.String(),
		},
	})
	return jwtToken.SignedString([]byte(jwtSecret))
}

func ValidateAccessToken(tokenString string) (id string, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method in auth token")
		}
		jwtSecret, _ := os.LookupEnv("JWT_SECRET")
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*AccessTokenClaims)
	if !ok || !token.Valid || claims.Subject == "" || claims.Type != utils_constants.ACCESS_TOKEN_TYPE || claims.Issuer != utils_constants.ISSUER {
		return "", errors.New("invalid token: authentication failed")
	}
	return claims.Subject, nil
}
