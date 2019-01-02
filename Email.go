package mail

import (
	"net/smtp"
	"strings"
)

// Email represents an email.
type Email struct {
	Credentials *EmailCredentials `json:"credentials"`
	Data        *EmailData        `json:"data"`
}

// Send sends the email.
func (e *Email) Send() error {
	address := strings.Join([]string{e.Credentials.Hostname, ":", e.Credentials.Port}, "")
	creds := smtp.PlainAuth("", e.Credentials.Address, e.Credentials.Password, e.Credentials.Hostname)
	return smtp.SendMail(address, creds, e.Data.From.Address, e.Data.GetRecipients(), e.Data.ToByteArray())
}
