package test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	email "github.com/FetchWeb/Email"
)

// TestEmail sends test emails based on the JSON files provided.
func TestEmail(t *testing.T) {
	creds := &email.Credentials{}
	data := &email.Data{}

	// Read test email credentials and unmarshal.
	dat, err := ioutil.ReadFile("test-credentials.json")
	if err != nil {
		t.Fatalf("Failed to read file %v", err)
		return
	}
	if err := json.Unmarshal(dat, &creds); err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
		return
	}

	// Read test email data and unmarshal.
	dat, err = ioutil.ReadFile("test-data.json")
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
		return
	}
	if err := json.Unmarshal(dat, &data); err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
		return
	}

	// Add attachments from file.
	if err := data.AddAttachmentFromFile("test-attachment-1.txt", false); err != nil {
		t.Fatalf("Failed to add attachment")
		return
	}
	if err := data.AddAttachmentFromFile("test-attachment-2.png", false); err != nil {
		t.Fatalf("Failed to add attachment")
		return
	}

	// Add attachments from buffer.
	dat, err = ioutil.ReadFile("test-attachment-3.pdf")
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
		return
	}
	data.AddAttachmentFromBuffer("test-attachment-3.pdf", dat, false)

	// Read email body template.
	dat, err = ioutil.ReadFile("test-template.html")
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
		return
	}
	data.Body = string(dat)

	// Send email.
	if err := email.Send(creds, data); err != nil {
		t.Fatalf("Failed to send email: %v", err)
	}
}
