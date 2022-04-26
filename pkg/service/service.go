package service

import (
	"Run_Hse_Run/pkg/emailer"
	"Run_Hse_Run/pkg/repository"
)

type Sender interface {
	SendEmail(email string, text string) error
}

type Service struct {
	Sender
}

func NewService(repo *repository.Repository, sender *emailer.EmailSender) *Service {
	return &Service{
		Sender: NewSenderService(sender),
	}
}
