package mail

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"mime"
	"net/mail"
	"path/filepath"
	"strings"
	"time"
)

// Message type constants.
const (
	MessageTypePlain = "text/plain"
	MessageTypeHTML  = "text/html"
)

// EmailData represents a smtp message.
type EmailData struct {
	From            mail.Address                `json:"from"`
	To              []string                    `json:"to"`
	CC              []string                    `json:"cc"`
	BCC             []string                    `json:"bcc"`
	ReplyTo         string                      `json:"replyto"`
	Subject         string                      `json:"subject"`
	Body            string                      `json:"body"`
	BodyContentType string                      `json:"bodycontenttype"`
	Headers         []Header                    `json:"headers"`
	Attachments     map[string]*EmailAttachment `json:"attachments"`
}

// AddAttachmentFromFile adds an attachment from a directory to the EmailData.
func (ed *EmailData) AddAttachmentFromFile(file string, inline bool) error {
	// Read the file.
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	// Get the filename.
	_, filename := filepath.Split(file)

	// Add file to attachments.
	ed.Attachments[filename] = &EmailAttachment{
		Filename: filename,
		Data:     data,
		Inline:   inline,
	}

	return nil
}

// AddAttachmentFromBuffer adds an attachment already in a byte array to the EmailData.
func (ed *EmailData) AddAttachmentFromBuffer(filename string, buffer []byte, inline bool) error {
	ed.Attachments[filename] = &EmailAttachment{
		Filename: filename,
		Data:     buffer,
		Inline:   inline,
	}
	return nil
}

// AddHeader ads a Header to the EmailData.
func (ed *EmailData) AddHeader(key string, value string) Header {
	newHeader := Header{Key: key, Value: value}
	ed.Headers = append(ed.Headers, newHeader)
	return newHeader
}

// GetRecipients returns all the recipients from the EmailData.
func (ed *EmailData) GetRecipients() []string {
	recipients := ed.To
	recipients = append(recipients, ed.CC...)
	recipients = append(recipients, ed.BCC...)
	return recipients
}

// ToByteArray returns all the EmailData as a byte array.
func (ed *EmailData) ToByteArray() []byte {
	buf := bytes.NewBuffer(nil)

	// Write from and recipients.
	buf.WriteString("From: " + ed.From.String() + "\n")
	buf.WriteString("Date: " + time.Now().Format(time.RFC1123Z) + "\n")
	buf.WriteString("To: " + strings.Join(ed.To, ",") + "\n")
	if len(ed.CC) > 0 {
		buf.WriteString("CC: " + strings.Join(ed.CC, ",") + "\n")
	}

	// Write encoding.
	var coder = base64.StdEncoding
	var subject = "=?UTF-8?B?" + coder.EncodeToString([]byte(ed.Subject)) + "?="
	buf.WriteString("Subject: " + subject + "\n")

	if len(ed.ReplyTo) > 0 {
		buf.WriteString("Reply-To: " + ed.ReplyTo + "\n")
	}

	buf.WriteString("MIME-Version: 1.0\n")

	// Write headers.
	if len(ed.Headers) > 0 {
		for _, header := range ed.Headers {
			buf.WriteString(fmt.Sprintf("%s: %s\n", header.Key, header.Value))
		}
	}

	// Write boundary.
	boundary := "f46d043c813270fc6b04c2d223da"
	if len(ed.Attachments) > 0 {
		buf.WriteString("Content-Type: multipart/mixed; boundary=" + boundary + "\n")
		buf.WriteString("\n--" + boundary + "\n")
	}

	// Write content type.
	buf.WriteString(fmt.Sprintf("Content-Type: %s; charset=utf-8\n\n", ed.BodyContentType))
	buf.WriteString(ed.Body)
	buf.WriteString("\n")

	// Write attachments.
	if len(ed.Attachments) > 0 {
		for _, attachment := range ed.Attachments {
			buf.WriteString("\n\n--" + boundary + "\n")

			if attachment.Inline {
				buf.WriteString("Content-Type: message/rfc822\n")
				buf.WriteString("Content-Disposition: inline; filename=\"" + attachment.Filename + "\"\n\n")

				buf.Write(attachment.Data)
			} else {
				ext := filepath.Ext(attachment.Filename)
				mimetype := mime.TypeByExtension(ext)
				if mimetype != "" {
					mime := fmt.Sprintf("Content-Type: %s\n", mimetype)
					buf.WriteString(mime)
				} else {
					buf.WriteString("Content-Type: application/octet-stream\n")
				}
				buf.WriteString("Content-Transfer-Encoding: base64\n")

				buf.WriteString("Content-Disposition: attachment; filename=\"=?UTF-8?B?")
				buf.WriteString(coder.EncodeToString([]byte(attachment.Filename)))
				buf.WriteString("?=\"\n\n")

				b := make([]byte, base64.StdEncoding.EncodedLen(len(attachment.Data)))
				base64.StdEncoding.Encode(b, attachment.Data)

				for i, l := 0, len(b); i < l; i++ {
					buf.WriteByte(b[i])
					if (i+1)%76 == 0 {
						buf.WriteString("\n")
					}
				}
			}

			buf.WriteString("\n--" + boundary)
		}

		buf.WriteString("--")
	}

	return buf.Bytes()
}
