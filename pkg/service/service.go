package service

import (
	"Run_Hse_Run/pkg/mailer"
	"Run_Hse_Run/pkg/repository"
	"sync"
)

var (
	mu    sync.Mutex
	Codes = make(map[string]int)
)

type Sender interface {
	SendEmail(email string) error
}

type Service struct {
	Sender
}

func NewService(repo *repository.Repository, sender *mailer.EmailSender) *Service {
	return &Service{
		Sender: NewSenderService(sender),
	}
}
