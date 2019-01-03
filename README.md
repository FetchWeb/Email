# FetchWeb Mail - A Simple SMTP Email API Written in Go

## Setup Example
```
package main

import (
	"encoding/json"
	"io/ioutil"

	mail "github.com/FetchWeb/Mail"
)

func main() {
	// Initailise credentials & data objects.
	emailCreds := &mail.EmailCredentials{
		Address: "<Sending email address>",
		Hostname: "<Hostname of SMTP Server>",
		Name: "<Name appearing on email>",
		Port: "<Port email being sent on>",
		Password "<Password to email account>"
	}
	emailData := &mail.EmailData{
		Attachments: make(map[string]*mail.EmailAttachment)
	}

	// Add email body.
	emailData.Body = "Hello world from FetchWeb Mail!"

	// Add attachment from file.
	if err := emailData.AddAttachmentFromFile("<Attachment directory>", false); err != nil {
		panic(err)
	}

	// Prepare and send email.
	email := &mail.Email{
		Credentials: emailCreds,
		Data: emailData
	}
	email.Send()
}

```
