package service

import (
	"errors"
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

func (a *AuthService) ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(os.Getenv("SIGNING_KEY")), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaim)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return "", errors.New("token expired")
	}

	return claims.Email, nil
}
