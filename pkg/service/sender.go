package service

import (
	"Run_Hse_Run/pkg/mailer"
	"bytes"
	"html/template"
	"log"
	"math/rand"
	"time"
)

type SenderService struct {
	sender *mailer.Mailer
}

func (s *SenderService) SendEmail(email string) error {
	code := rand.Intn(9000) + 1000

	Mu.Lock()
	Codes[email] = code
	Mu.Unlock()

	buffer := bytes.NewBufferString("")
	tmpl, _ := template.ParseFiles("templates/message.html")
	err := tmpl.Execute(buffer, struct {
		Code int
	}{
		Code: code,
	})

	if err != nil {
		log.Fatalf("Can't read template file: %s", err.Error())
	}

	return s.sender.SendEmail(email, buffer.String())
}

func NewSenderService(sender *mailer.Mailer) *SenderService {
	rand.Seed(time.Now().Unix())
	return &SenderService{sender: sender}
}
