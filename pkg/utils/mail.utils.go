package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"net/mail"
	"net/smtp"
	"strings"
)

var (
	user     = MustGetEnv("MAIL_USER")
	password = MustGetEnv("MAIL_PASSWORD")
	from     = MustGetEnv("MAIL_FROM")
	fromName = GetEnvFallback("MAIL_FROM_NAME", from)
	smtpHost = MustGetEnv("MAIL_SMTP")
	smtpPort = MustGetEnv("MAIL_SMTP_PORT")
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

func encodeRFC2047(str string) string {
	// use mail's rfc2047 to encode any string
	addr := mail.Address{Name: str, Address: ""}
	return strings.Trim(addr.String(), " <>")
}

func SendMailTemplate(info MailWithTemplate) error {
	auth := smtp.PlainAuth("", user, password, smtpHost)

	from := mail.Address{Name: fromName, Address: from}

	t, _ := template.ParseFiles(info.Template)

	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["Subject"] = encodeRFC2047(info.Subject)
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
