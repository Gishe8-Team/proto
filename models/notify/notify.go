package notify

type SmsModel struct {
	MessagesText  string
	MobileNumbers []string
}

type EmailModel struct {
	Message        string
	EmailAddresses []string
}
