package mail

import (
	"bytes"
	"fmt"
	"net/mail"
	"net/smtp"
	"text/template"
)

var (
	user     = "MAIL_USER"
	password = "MAIL_PASSWORD"
	from     = "MAIL_FROM"
	fromName = "MAIL_FROM_NAME"
	smtpHost = "MAIL_SMTP"
	smtpPort = "MAIL_SMTP_PORT"
)

type Mail struct {
	To      []string
	Subject string
}

type MailWithTemplate struct {
	Mail
	Template string
	Data     any
}

func SendMailTemplate(info MailWithTemplate) error {
	auth := smtp.PlainAuth("", user, password, smtpHost)

	from := mail.Address{Name: fromName, Address: from}

	t, _ := template.ParseFiles(info.Template)

	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["Subject"] = info.Subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/html; charset=\"UTF-8\""

	var body bytes.Buffer
	for k, v := range headers {
		body.Write([]byte(fmt.Sprintf("%s: %s\r\n", k, v)))
	}

	t.Execute(&body, info.Data)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from.Address, info.To, body.Bytes())
	if err != nil {
		return err
	}
	return nil
}
