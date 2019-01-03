package test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	mail "github.com/FetchWeb/Mail"
)

// TestEmail sends test emails based on the JSON files provided.
func TestEmail(t *testing.T) {
	emailCreds := &mail.EmailCredentials{}
	emailData := &mail.EmailData{Attachments: make(map[string]*mail.EmailAttachment)}

	// Read test email credentials and unmarshal.
	dat, err := ioutil.ReadFile("TestEmailCredentials.json")
	if err != nil {
		t.Errorf("ERROR - Failed to read in TestEmailCredentials.json: %v", err)
		return
	}
	if err := json.Unmarshal(dat, &emailCreds); err != nil {
		t.Errorf("ERROR - Failed to unmarshal JSON data: %v", err)
		return
	}

	// Read test email data and unmarshal.
	dat, err = ioutil.ReadFile("TestEmailData.json")
	if err != nil {
		t.Errorf("ERROR - Failed to read in TestEmailData.json: %v", err)
		return
	}
	if err := json.Unmarshal(dat, &emailData); err != nil {
		t.Errorf("ERROR - Failed to unmarshal JSON data: %v", err)
		return
	}

	// Add attachments from file.
	if err := emailData.AddAttachmentFromFile("TestEmailAttachment_1.txt", false); err != nil {
		t.Errorf("ERROR - Failed to add attachment 1")
		return
	}
	if err := emailData.AddAttachmentFromFile("TestEmailAttachment_2.png", false); err != nil {
		t.Errorf("ERROR - Failed to add attachment 2")
		return
	}

	// Add attachments from buffer.
	dat, err = ioutil.ReadFile("TestEmailAttachment_3.pdf")
	if err != nil {
		t.Errorf("ERROR - Failed to read in TestEmailAttachment_3.pdf: %v", err)
		return
	}
	if err := emailData.AddAttachmentFromBuffer("TestEmailAttachment_3.pdf", dat, false); err != nil {
		t.Errorf("ERROR - Failed to add attachment 3")
		return
	}

	// Read email body template.
	dat, err = ioutil.ReadFile("TestEmailTemplate.html")
	if err != nil {
		t.Errorf("ERROR - Failed to read in TestEmailTemplate.html: %v", err)
		return
	}
	emailData.Body = string(dat)

	// Prepare and send email.
	email := &mail.Email{Credentials: emailCreds, Data: emailData}
	email.Send()
}
