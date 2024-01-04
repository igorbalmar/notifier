package sendmail

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"text/template"
)

func SendMail(recipients []string, subject string, server string, errMsg string, timestamp string) {
	var from string = "igbalmar@gmail.com"
	var pass string = "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"
	var body bytes.Buffer
	mimeHeader := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	smtpHost := "smtp.sendgrid.net"
	smtpPort := "587"
	t, _ := template.ParseFiles("template.html")

	if len(os.Args) < 2 {
		fmt.Println("This program is meant to be used to send e-mails to the address informed as parameters from OS")
		os.Exit(1)
	}

	body.Write([]byte(fmt.Sprintf("From: %s\nSubject: %s\n%s\n\n", from, subject, mimeHeader)))
	t.Execute(&body, struct {
		Server  string
		Error   string
		Horario string
	}{
		Server:  server,
		Error:   errMsg,
		Horario: timestamp,
	})

	err := smtp.SendMail(smtpHost+":"+smtpPort,
		smtp.PlainAuth("", "apikey", pass, smtpHost),
		from,
		recipients,
		body.Bytes())
	if err != nil {
		log.Fatal(err)
	}

}
