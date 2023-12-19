package main

import (
	"fmt"

	"github.com/psanti93/galleryValleyv1/models"
)

const (
	host     = "sandbox.smtp.mailtrap.io"
	port     = 2525
	username = "d26bb25ad79d5b"
	password = "c3b937b646f1d3"
)

func main() {

	to := "paul@test.com"
	from := "go@gest.com"
	subject := "Hello World"
	plainText := "Hello World! TESTING 123"
	html := `<h1>Hello World!</h1><p>This is a test email</p>`

	email := models.Email{
		From:      from,
		To:        to,
		Subject:   subject,
		Plaintext: plainText,
		HTML:      html,
	}

	es := models.NewEmailService(models.SMTConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	})

	err := es.Send(email)

	if err != nil {
		panic(err)
	}
	fmt.Println("Message sent!")
}
