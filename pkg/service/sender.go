package service

import "Run_Hse_Run/pkg/emailer"

type SenderService struct {
	sender *emailer.EmailSender
}

func (s *SenderService) SendEmail(email string, text string) error {
	return s.SendEmail(email, text)
}

func NewSenderService(sender *emailer.EmailSender) *SenderService {
	return &SenderService{sender: sender}
}
