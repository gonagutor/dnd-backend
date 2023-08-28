package auth_utils

import (
	utils_constants "dnd/backend/constants/utils"
	"dnd/backend/models"
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type RefreshTokenClaims struct {
	CustomKey string `json:"key"`
	Type      string `json:"type"`
	jwt.RegisteredClaims
}

func GenerateRefreshToken(user *models.User) (string, error) {
	jwtSecret, _ := os.LookupEnv("JWT_SECRET")
	refreshKey, err := bcrypt.GenerateFromPassword([]byte(user.ID.String()+user.RefreshKey), utils_constants.REFRESH_TOKEN_COST)
	if err != nil {
		return string(refreshKey), err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS512, RefreshTokenClaims{
		string(refreshKey),
		utils_constants.REFRESH_TOKEN_TYPE,
		jwt.RegisteredClaims{
			Issuer:  utils_constants.ISSUER,
			Subject: user.ID.String(),
		},
	})
	return jwtToken.SignedString([]byte(jwtSecret))
}

func ValidateRefreshToken(tokenString string) (id string, key string, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &RefreshTokenClaims{}, KeyFunc)
	if err != nil {
		return "", "", err
	}

	claims, ok := token.Claims.(*RefreshTokenClaims)
	if !ok || !token.Valid || claims.Subject == "" || claims.Type != utils_constants.REFRESH_TOKEN_TYPE || claims.Issuer != utils_constants.ISSUER {
		return "", "", errors.New("invalid token: authentication failed")
	}
	return claims.Subject, claims.CustomKey, nil
}
