package service

import (
	"Run_Hse_Run/pkg/mailer"
	"Run_Hse_Run/pkg/model"
	"Run_Hse_Run/pkg/repository"
	"sync"
)

var (
	Mu    sync.Mutex
	Codes = make(map[string]int)
)

type Sender interface {
	SendEmail(email string) error
}

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(email string) (model.User, error)
	GenerateToken(email string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Service struct {
	Sender
	Authorization
}

func NewService(repo *repository.Repository, sender *mailer.Mailer) *Service {
	return &Service{
		Sender:        NewSenderService(sender),
		Authorization: NewAuthService(repo),
	}
}
