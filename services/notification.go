package services

const (
	SendAnSMS         = "15.1.1"
	SendBulkSMS       = "15.1.2"
	GetAccountBalance = "15.1.3"
	SendCodeSMS       = "15.1.4"
)

const (
	SendAnEmail   = "15.2.1"
	SendBulkEmail = "15.2.2"
	SendCodeEmail = "15.2.3"
)

const (
	SendPushNotification = "15.3.1"
	SendWebNotification  = "15.3.2"
)

const (
	ListSMSArchive          = "15.4.1"
	FindSMSArchiveByMobile  = "15.4.2"
	FindSMSArchiveByService = "15.4.3"

	ListEmailArchive               = "15.4.4"
	FindEmailArchiveByEmailAddress = "15.4.5"
	FindEmailArchiveByService      = "15.4.6"
)
