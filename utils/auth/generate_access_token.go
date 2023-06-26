package auth_utils

import (
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
