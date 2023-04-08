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

type SMSArchiveRequest struct {
	UserID       string `json:"user_id,omitempty"`
	MobileNumber string `json:"mobile_number,omitempty"`
	Service      string `json:"service,omitempty"`
	Limit        int    `json:"limit,omitempty"`
	Offset       int    `json:"offset,omitempty"`
}

type SMSArchiveSlice struct {
	ID           string `json:"id" boil:"id"`
	MobileNum    string `json:"mobile_num" boil:"mobile_num"`
	Message      string `json:"message" boil:"message"`
	WhichService string `json:"which_service" boil:"which_service"`
}

type EmailArchiveRequest struct {
	UserID       string `json:"user_id,omitempty"`
	EmailAddress string `json:"email_address,omitempty"`
	Service      string `json:"service,omitempty"`
	Limit        int    `json:"limit,omitempty"`
	Offset       int    `json:"offset,omitempty"`
}

type EmailArchiveSlice struct {
	ID           string `json:"id" boil:"id"`
	EmailAddress string `json:"email_address" boil:"email_address"`
	EmailMessage string `json:"email_message" boil:"email_message"`
	WhichService string `json:"which_service" boil:"which_service"`
}
