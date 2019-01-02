package mail

// EmailCredentials stores the relevant credenial data to send emails.
type EmailCredentials struct {
	Address  string `json:"address"`
	Hostname string `json:"hostname"`
	Name     string `json:"name"`
	Port     string `json:"port"`
	Password string `json:"password"`
}
