package email

// Attachment represents an email attachment.
type Attachment struct {
	Filename string `json:"filename"`
	Data     string `json:"data"`
	Inline   bool   `json:"inline"`
}
