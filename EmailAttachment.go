package mail

// EmailAttachment represents an email attachment.
type EmailAttachment struct {
	Filename string `json:"filename"`
	Data     []byte `json:"data"`
	Inline   bool   `json:"inline"`
}
