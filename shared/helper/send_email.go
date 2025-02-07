package helper

import (
	"crypto/tls"
	"log"
	"os"
	"strconv"

	"gopkg.in/mail.v2"
)

func SendEmail(subject, body string, recipient ...string) error {

	if len(recipient) == 0 {
		return nil
	}

	notEmptyRecipient := make([]string, 0)
	for _, r := range recipient {
		if r != "" {
			notEmptyRecipient = append(notEmptyRecipient, r)
		}
	}

	if len(notEmptyRecipient) == 0 {
		return nil
	}

	from := os.Getenv("GMAIL_FROM")
	portStr := os.Getenv("GMAIL_SMTP_PORT")
	host := os.Getenv("GMAIL_SMTP_HOST")
	user := os.Getenv("GMAIL_USERNAME")
	pass := os.Getenv("GMAIL_APP_PASSWORD")

	// Create a new message
	m := mail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", notEmptyRecipient...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return err
	}

	// Create a new SMTP client
	d := mail.NewDialer(host, port, user, pass)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send the email
	if err := d.DialAndSend(m); err != nil {

		log.Printf("emails : %v", recipient)

		return err
	}

	return nil
}
