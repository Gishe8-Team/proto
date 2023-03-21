package notify

type SmsModel struct {
	MessagesText  string   `json:"messages_text"`
	MobileNumbers []string `json:"mobile_numbers"`
}

type EmailModel struct {
	Message        string   `json:"message"`
	Subject        string   `json:"subject"`
	EmailAddresses []string `json:"email-addresses"`
}
