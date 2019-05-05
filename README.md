# FetchWeb Mail

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/FetchWeb/Email)](https://goreportcard.com/report/github.com/FetchWeb/Email)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/FetchWeb/Email)
[![GitHub release](https://img.shields.io/github/release/FetchWeb/Email.svg)](https://github.com/FetchWeb/Email/releases )

## Introduction
FetchWeb Email is a Simple SMTP Email API written in Go with no dependencies outside of the standard library.

## Setup Example
```go
package main

import (
	"encoding/json"
	"io/ioutil"

	email "github.com/FetchWeb/Email"
)

func main() {
	// Initailise credentials & data objects.
	creds := &email.Credentials{
		Address: "<Sending email address>",
		Hostname: "<Hostname of SMTP Server>",
		Name: "<Name appearing on email>",
		Port: "<Port email being sent on>",
		Password "<Password to email account>"
	}
	data := &email.Data{
		Attachments: make(map[string]*email.Attachment)
	}

	// Add email body.
	emailData.Body = "Hello world from FetchWeb Mail!"

	// Add attachment from file.
	if err := emailData.AddAttachmentFromFile("<Attachment directory>", false); err != nil {
		panic(err)
	}

	// Send email.
	if err := email.Send(creds, data); err != nil {
		panic(err)
	}
}
```
