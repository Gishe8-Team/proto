package notify

type EmailModel struct {
	Message        string `json:"message" boil:"email_message"`
	Subject        string `json:"subject" boil:"email_subject"`
	EmailAddresses string `json:"email_addresses" boil:"email_address"`
	Service        string `json:"service,omitempty" boil:"which_service"`
}

type CodeEmailModel struct {
	EmailModel
	CodeToken string `json:"code_token"`
}

type BulkEmailModel struct {
	Message        string   `json:"message"`
	Subject        string   `json:"subject"`
	EmailAddresses []string `json:"email-addresses"`
	Service        string   `json:"service,omitempty"`
}
