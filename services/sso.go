package services

const (
	UserLogin     = "2.1.1"
	OTPLogin      = "2.1.2"
	EditLogin     = "2.1.3"
	CreateAccount = "2.1.4"
	DeleteAccount = "2.1.5"
	BanAccount    = "2.1.6"
	UnbanAccount  = "2.1.7"
	VerifyEmail   = "2.1.8"
	VerifyMobile  = "2.1.9"
)

const (
	CreateACL   = "2.2.1"
	EditACL     = "2.2.2"
	FindACLByID = "2.2.3"
	ListACL     = "2.2.4"
)

const (
	CreateUserGroup       = "2.3.1"
	EditUserGroup         = "2.3.2"
	AssignToUserGroup     = "2.3.3"
	CreateACLForUserGroup = "2.3.4"
	EditACLForUserGroup   = "2.3.5"
)

const (
	AddLoginHistoryEntry    = "2.4.1"
	ListLoginHistoryForUser = "2.4.2"
)
