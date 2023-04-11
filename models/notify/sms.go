package notify

type SmsModel struct {
	MobileNum    string `boil:"mobile_num" json:"mobile_num" toml:"mobile_num" yaml:"mobile_num"`
	Message      string `boil:"message" json:"message" toml:"message" yaml:"message"`
	WhichService string `boil:"which_service" json:"which_service" toml:"which_service" yaml:"which_service"`
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
