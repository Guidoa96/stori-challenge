package email

import (
	"fmt"
	"net/smtp"
)

func SendEmail(to, subject, body string) error {
	from := "garribas@fi.uba.ar"
	password := "Joseputo1"

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	msg := fmt.Sprintf("Subject: %s\n\n%s", subject, body)
	auth := smtp.PlainAuth("", from, password, smtpHost)

	return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(msg))
}
