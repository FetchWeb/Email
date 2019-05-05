# FetchWeb Mail

## Introduction
FetchWeb Mail is a Simple SMTP Email API written in Go with no dependencies outside of the standard library.

## Setup Example
```go
package main

import (
	"encoding/json"
	"io/ioutil"

	email "github.com/FetchWeb/Mail"
)

func main() {
	// Initailise credentials & data objects.
	emailCreds := &email.Credentials{
		Address: "<Sending email address>",
		Hostname: "<Hostname of SMTP Server>",
		Name: "<Name appearing on email>",
		Port: "<Port email being sent on>",
		Password "<Password to email account>"
	}
	emailData := &email.Data{
		Attachments: make(map[string]*email.Attachment)
	}

	// Add email body.
	emailData.Body = "Hello world from FetchWeb Mail!"

	// Add attachment from file.
	if err := emailData.AddAttachmentFromFile("<Attachment directory>", false); err != nil {
		panic(err)
	}

	// Prepare and send email.
	email := &email.Email{
		Credentials: emailCreds,
		Data: emailData
	}
	email.Send()
}
```
