package auth_utils

import (
	"os"
	utils_constants "revosearch/backend/constants/utils"
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
