package services

const (
	CreateAccount = "2.1.1"
	EditAccount   = "2.1.2"
	DeleteAccount = "2.1.3"
	BanAccount    = "2.1.4"
	UnbanAccount  = "2.1.5"
	VerifyEmail   = "2.1.6"
	VerifyMobile  = "2.1.7"
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
	UserLogin = "2.4.1"
	OTPLogin  = "2.4.2"
)

const (
	AddLoginHistoryEntry    = "2.5.1"
	ListLoginHistoryForUser = "2.5.2"
)
