package models

import (
	"fmt"

	"github.com/go-mail/mail/v2"
)

type Email struct {
	From      string
	To        string
	Subject   string
	Plaintext string
	HTML      string
}

type SMTConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

const (
	DefaultSender = "support@lenslocked.com"
)

type EmailService struct {
	DefaultSender string
	//unexported fields

	dailer *mail.Dialer
}

func NewEmailService(config SMTConfig) *EmailService {
	es := EmailService{
		dailer: mail.NewDialer(config.Host, config.Port, config.Username, config.Password),
	}

	return &es
}

func (es *EmailService) Send(email Email) error {
	msg := mail.NewMessage()
	msg.SetHeader("To", email.To)
	es.setFrom(msg, email) // Set FROM field to a default value if it's not set in the email
	msg.SetHeader("Subject", email.Subject)
	switch {
	case email.Plaintext != "" && email.HTML != "":
		msg.SetBody("text/plain", email.Plaintext)
		msg.AddAlternative("text/html", email.HTML)
	case email.Plaintext != "":
		msg.SetBody("text/plain", email.Plaintext)
	case email.HTML != "":
		msg.SetBody("text/html", email.HTML)
	}

	err := es.dailer.DialAndSend(msg)
	if err != nil {
		return fmt.Errorf("send: %w", err)
	}
	return nil
}

func (es *EmailService) ForgotPassword(to, resetURL string) error {
	email := Email{
		Subject:   "Reset our Password",
		To:        to,
		Plaintext: "To reset your password please visit the following link: " + resetURL,
		HTML: `<p>To Reset your password, please visit the following link:<a href="` + resetURL + `">` +
			resetURL + `</a></p>`,
	}
	err := es.Send(email)
	if err != nil {
		return fmt.Errorf("Forgot Password Email: %w", err)
	}
	return nil
}

func (es EmailService) setFrom(msg *mail.Message, email Email) {
	var from string
	switch {
	case email.From != "":
		from = email.From
	case es.DefaultSender != "":
		from = es.DefaultSender
	default:
		from = DefaultSender
	}

	msg.SetHeader("From", from)
}
