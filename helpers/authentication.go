package helpers

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kidboy-man/mini-bank-rest/configs"
	"github.com/kidboy-man/mini-bank-rest/middlewares"
	"github.com/kidboy-man/mini-bank-rest/models"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 16)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateToken(user *models.User) (result string, err error) {
	claims := middlewares.JWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "mini-bank-rest",
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(configs.AppConfig.JWTExpirationSecond)},
		},
		UserID:   user.ID,
		Username: user.Username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	result, err = token.SignedString([]byte(configs.AppConfig.JWTSignatureKey))
	return
}
