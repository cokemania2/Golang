package main

import (
	"fmt"
	"net/smtp"
)

func main() {
	// some configuration
	hostURL := "smtp.gmail.com"
	hostPort := "587"
	emailSender := SENDER
	password := PASSWORD
	emailRecevier := RECEVIER

	// Create Auth
	emailAuth := smtp.PlainAuth(
		"",
		emailSender,
		password,
		hostURL,
	)

	// Create Email
	msg := []byte("To : " + emailRecevier + "\r\n" +
		"Subject: " + "Hello" + "\r\n" + "How are you doing")

	// Send Mail
	err := smtp.SendMail(
		hostURL+":"+hostPort,
		emailAuth,
		emailSender,
		[]string{emailRecevier},
		msg)
	if err != nil {
		fmt.Print("Error :", err)
	}
	fmt.Print("Email sent")
}
