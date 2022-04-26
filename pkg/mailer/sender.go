package mailer

import (
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

type SenderEmail struct {
	dialer *gomail.Dialer
}

func NewSendEmail(dialer *gomail.Dialer) *SenderEmail {
	return &SenderEmail{dialer: dialer}
}

func (s *SenderEmail) SendEmail(email string, text string) error {
	message := gomail.NewMessage()
	message.SetHeader("From", viper.GetString("mailer.email"))
	message.SetHeader("To", email)
	message.SetHeader("Subject", "Authorization in \"Run, Hse, Run\"")
	message.SetBody("text/html", text)

	return s.dialer.DialAndSend(message)
}
