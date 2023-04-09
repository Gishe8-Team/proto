package notify

type SmsModel struct {
	MessagesText  string `json:"messages_text"`
	MobileNumbers string `json:"mobile_numbers"`
	Service       string `json:"service,omitempty"`
}

type BulkSmsModel struct {
	MessagesText  string   `json:"messages_text"`
	MobileNumbers []string `json:"mobile_numbers"`
	Service       string   `json:"service,omitempty"`
}

type Redirect struct {
	RedirectTo string `json:"redirect_to"`
}
