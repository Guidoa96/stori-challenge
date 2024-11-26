package email

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendEmail(to, subject, body string) error {
	// Read environment variables
	smtpServer := os.Getenv("SMTP_SERVER")
	smtpPort := os.Getenv("SMTP_PORT")
	emailFrom := os.Getenv("EMAIL_FROM")

	// Ensure environment variables are set
	if smtpServer == "" || smtpPort == "" || emailFrom == "" {
		return fmt.Errorf("SMTP configuration not set")
	}

	message := []byte(fmt.Sprintf(
		"Subject: %s\r\n"+
			"MIME-Version: 1.0\r\n"+
			"Content-Type: text/html; charset=\"UTF-8\"\r\n"+
			"\r\n"+
			"%s",
		subject, body))

	err := smtp.SendMail(fmt.Sprintf("%s:%s", smtpServer, smtpPort), nil, emailFrom, []string{to}, message)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	fmt.Println("Email sent successfully!")
	return nil
}
