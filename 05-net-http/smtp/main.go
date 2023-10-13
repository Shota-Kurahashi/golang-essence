package main

import (
	"log"
	"net/smtp"

	"github.com/jhillyerd/enmime"
)

func main() {
	smtpHost := "smtp.gmail.com"
	smtpAuth := smtp.PlainAuth("example.com", "example-user", "example-password", "auth.example.com")

	sender := enmime.NewSMTP(smtpHost, smtpAuth)

	master := enmime.Builder().From("送信太郎", "taro@example.com").Subject("テストメール").Text([]byte("テストメールです。")).HTML([]byte("<h1>テストメールです。</h1>"))

	msg := master.To("受信花子", "hanako@example.com")

	err := msg.Send(sender)

	if err != nil {
		log.Fatal(err)
	}
}
