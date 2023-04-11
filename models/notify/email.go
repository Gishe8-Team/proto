package notify

import "github.com/volatiletech/null/v8"

type EmailModel struct {
	EmailAddress string      `boil:"email_address" json:"email_address" toml:"email_address" yaml:"email_address"`
	EmailMessage string      `boil:"email_message" json:"email_message" toml:"email_message" yaml:"email_message"`
	WhichService string      `boil:"which_service" json:"which_service" toml:"which_service" yaml:"which_service"`
	EmailSubject null.String `boil:"email_subject" json:"email_subject,omitempty" toml:"email_subject" yaml:"email_subject,omitempty"`
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
