package mailer

import "gopkg.in/gomail.v2"

type Sender interface {
	SendEmail(email string, text string) error
}

type Mailer struct {
	Sender
}

func NewMailer(dialer *gomail.Dialer) *Mailer {
	return &Mailer{
		Sender: NewSendEmail(dialer),
	}
}
