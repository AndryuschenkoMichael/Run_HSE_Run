package service

import (
	"Run_Hse_Run/pkg/model"
	"Run_Hse_Run/pkg/repository"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

const (
	tokenTTL = 24 * 365 * 2 * time.Hour
)

type AuthService struct {
	repo *repository.Repository
}

func (a *AuthService) CreateUser(user model.User) (int, error) {
	return a.repo.CreateUser(user)
}

func (a *AuthService) GetUser(email string) (model.User, error) {
	return a.repo.GetUser(email)
}

func NewAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}

type tokenClaim struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func (a *AuthService) GenerateToken(email string) (string, error) {
	user, err := a.GetUser(email)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaim{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(os.Getenv("SIGNING_KEY")))
}

func (a *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(os.Getenv("SIGNING_KEY")), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaim)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return 0, errors.New("token expired")
	}

	return claims.UserId, nil
}
