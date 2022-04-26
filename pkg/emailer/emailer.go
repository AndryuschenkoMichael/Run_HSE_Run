package emailer

import "gopkg.in/gomail.v2"

type Sender interface {
	SendEmail(email string, text string) error
}

type EmailSender struct {
	Sender
}

func NewEmailSender(dialer *gomail.Dialer) *EmailSender {
	return &EmailSender{
		Sender: NewSendEmail(dialer),
	}
}
