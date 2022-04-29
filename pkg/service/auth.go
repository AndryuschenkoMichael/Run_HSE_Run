package service

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

const (
	tokenTTL = 24 * 365 * 2 * time.Hour
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

type tokenClaim struct {
	jwt.StandardClaims
	Email string `json:"email"`
}

func (a *AuthService) GenerateToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaim{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		email,
	})

	return token.SignedString([]byte(os.Getenv("SIGNING_KEY")))
}
