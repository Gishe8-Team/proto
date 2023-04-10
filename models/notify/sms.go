package notify

type SmsModel struct {
	MessagesText  string `json:"messages_text"`
	MobileNumbers string `json:"mobile_numbers"`
	Service       string `json:"service,omitempty"`
}

type CodeSmsModel struct {
	SmsModel
	CodeToken string `json:"code_token"`
}

type BulkSmsModel struct {
	MessagesText  string   `json:"messages_text"`
	MobileNumbers []string `json:"mobile_numbers"`
	Service       string   `json:"service,omitempty"`
}

type Redirect struct {
	RedirectTo string `json:"redirect_to"`
	CodeToken  string `json:"code_token"`
}
