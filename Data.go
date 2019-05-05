package email

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

// Data represents a smtp email.
type Data struct {
	From            mail.Address           `json:"from"`
	To              []string               `json:"to"`
	CC              []string               `json:"cc"`
	BCC             []string               `json:"bcc"`
	ReplyTo         string                 `json:"reply_to"`
	Subject         string                 `json:"subject"`
	Body            string                 `json:"body"`
	BodyContentType string                 `json:"body_content_type"`
	Headers         []Header               `json:"headers"`
	Attachments     map[string]*Attachment `json:"attachments"`
}

// AddAttachmentFromFile adds an attachment from a directory to the Data.
func (d *Data) AddAttachmentFromFile(file string, inline bool) error {
	// Read the file.
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	// Get the filename.
	_, filename := filepath.Split(file)

	// Add file to attachments.
	d.Attachments[filename] = &Attachment{
		Filename: filename,
		Data:     data,
		Inline:   inline,
	}

	return nil
}

// AddAttachmentFromBuffer adds an attachment already in a byte array to the Data.
func (d *Data) AddAttachmentFromBuffer(filename string, buffer []byte, inline bool) error {
	d.Attachments[filename] = &Attachment{
		Filename: filename,
		Data:     buffer,
		Inline:   inline,
	}
	return nil
}

// AddHeader ads a Header to the Data.
func (d *Data) AddHeader(key string, value string) Header {
	newHeader := Header{Key: key, Value: value}
	d.Headers = append(d.Headers, newHeader)
	return newHeader
}

// GetRecipients returns all the recipients from the Data.
func (d *Data) GetRecipients() []string {
	recipients := d.To
	recipients = append(recipients, d.CC...)
	recipients = append(recipients, d.BCC...)
	return recipients
}

// ToByteArray returns all the Data as a byte array.
func (d *Data) ToByteArray() []byte {
	buf := bytes.NewBuffer(nil)

	// Write from and recipients.
	buf.WriteString("From: " + d.From.String() + "\n")
	buf.WriteString("Date: " + time.Now().Format(time.RFC1123Z) + "\n")
	buf.WriteString("To: " + strings.Join(d.To, ",") + "\n")
	if len(d.CC) > 0 {
		buf.WriteString("CC: " + strings.Join(d.CC, ",") + "\n")
	}

	// Write encoding.
	var coder = base64.StdEncoding
	var subject = "=?UTF-8?B?" + coder.EncodeToString([]byte(d.Subject)) + "?="
	buf.WriteString("Subject: " + subject + "\n")

	if len(d.ReplyTo) > 0 {
		buf.WriteString("Reply-To: " + d.ReplyTo + "\n")
	}

	buf.WriteString("MIME-Version: 1.0\n")

	// Write headers.
	if len(d.Headers) > 0 {
		for _, header := range d.Headers {
			buf.WriteString(fmt.Sprintf("%s: %s\n", header.Key, header.Value))
		}
	}

	// Write boundary.
	boundary := "f46d043c813270fc6b04c2d223da"
	if len(d.Attachments) > 0 {
		buf.WriteString("Content-Type: multipart/mixed; boundary=" + boundary + "\n")
		buf.WriteString("\n--" + boundary + "\n")
	}

	// Write content type.
	buf.WriteString(fmt.Sprintf("Content-Type: %s; charset=utf-8\n\n", d.BodyContentType))
	buf.WriteString(d.Body)
	buf.WriteString("\n")

	// Write attachments.
	if len(d.Attachments) > 0 {
		for _, attachment := range d.Attachments {
			buf.WriteString("\n\n--" + boundary + "\n")

			if attachment.Inline {
				buf.WriteString("Content-Type: email/rfc822\n")
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
