package main

import (
	"fmt"
	"os"

	"github.com/go-mail/mail/v2"
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
	plainText := "Hello World!"
	html := `<h1>Hello World!</h1><p>This is a test email</p>`

	msg := mail.NewMessage()
	msg.SetHeader("To", to)
	msg.SetHeader("From", from)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/plain", plainText)
	msg.AddAlternative("text/html", html)

	msg.WriteTo(os.Stdout)

	dailer := mail.NewDialer(host, port, username, password)
	// One approach
	// sender, err := dailer.Dial()

	// if err != nil {
	// 	panic(err)
	// }

	// defer sender.Close()
	// sender.Send(from, []string{to}, msg)

	err := dailer.DialAndSend(msg)
	if err != nil {
		panic(err)
	}
	fmt.Println("Message sent!")
}
