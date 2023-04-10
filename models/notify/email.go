package notify

type EmailModel struct {
	Message        string `json:"message"`
	Subject        string `json:"subject"`
	EmailAddresses string `json:"email-addresses"`
	Service        string `json:"service,omitempty"`
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
