package models

import "github.com/go-mail/mail/v2"

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
