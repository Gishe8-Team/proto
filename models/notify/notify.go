package notify

type SmsModel struct {
	MessagesText  string   `json:"messages_text"`
	MobileNumbers []string `json:"mobile_numbers"`
	Service       string   `json:"service,omitempty"`
}

type EmailModel struct {
	Message        string   `json:"message"`
	Subject        string   `json:"subject"`
	EmailAddresses []string `json:"email-addresses"`
	Service        string   `json:"service,omitempty"`
}

type SMSArchiveSlice struct {
	ID        string `json:"id"`
	MobileNum string `json:"mobile_num"`
	Message   string `json:"message"`
	Service   string `json:"service"`
}

type EmailArchiveSlice struct {
	ID           string `json:"id"`
	EmailAddress string `json:"email_address"`
	Message      string `json:"message"`
	Service      string `json:"service"`
}
