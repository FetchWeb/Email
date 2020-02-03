package email

// Credentials stores the relevant credenial data to send emails.
type Credentials struct {
	Address  string `json:"address"`
	Hostname string `json:"hostname"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Port     string `json:"port"`
}
