package email

// Header represents an additional email header.
type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
