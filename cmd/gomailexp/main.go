package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/psanti93/galleryValleyv1/models"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("SMTP_HOST")
	portStr := os.Getenv("SMTP_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		panic(err)
	}
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")

	// to := "paul@test.com"
	// from := "go@gest.com"
	// subject := "Hello World"
	// plainText := "Hello World! TESTING 123"
	// html := `<h1>Hello World!</h1><p>This is a test email</p>`

	// email := models.Email{
	// 	From:      from,
	// 	To:        to,
	// 	Subject:   subject,
	// 	Plaintext: plainText,
	// 	HTML:      html,
	// }

	es := models.NewEmailService(models.SMTConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	})

	err = es.ForgotPassword("paul@test.com", "http://lenslockedv1.com/pw?token=1234")

	if err != nil {
		panic(err)
	}
	fmt.Println("Message sent!")
}
