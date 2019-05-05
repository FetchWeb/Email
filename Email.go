package email

import (
	"net/smtp"
	"strings"
)

// Send sends the email.
func Send(credentials *Credentials, data *Data) error {
	address := strings.Join([]string{credentials.Hostname, ":", credentials.Port}, "")
	creds := smtp.PlainAuth("", credentials.Address, credentials.Password, credentials.Hostname)
	return smtp.SendMail(address, creds, data.From.Address, data.GetRecipients(), data.ToByteArray())
}
