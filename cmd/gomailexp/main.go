package main

import (
	"os"

	"github.com/go-mail/mail/v2"
)

func main() {

	html := `<h1>Hello World!</h1><p>This is a test email</p>`

	msg := mail.NewMessage()
	msg.SetHeader("To", "paul@test.com")
	msg.SetHeader("From", "go@gest.com")
	msg.SetHeader("Subject", "Hello World")
	msg.SetBody("text/plain", "Hello World!")
	msg.AddAlternative("text/html", html)

	msg.WriteTo(os.Stdout)

}
