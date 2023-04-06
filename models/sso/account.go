package sso

import "github.com/volatiletech/null/v8"

type Account struct {
	ID             string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserName       null.String `boil:"user_name" json:"user_name,omitempty" toml:"user_name" yaml:"user_name,omitempty"`
	Password       null.String `boil:"password" json:"password,omitempty" toml:"password" yaml:"password,omitempty"`
	GroupID        int16       `boil:"group_id" json:"group_id" toml:"group_id" yaml:"group_id"`
	UserACL        null.String `boil:"user_acl" json:"user_acl,omitempty" toml:"user_acl" yaml:"user_acl,omitempty"`
	UserEmail      null.String `boil:"user_email" json:"user_email,omitempty" toml:"user_email" yaml:"user_email,omitempty"`
	EmailVerified  null.Bool   `boil:"email_verified" json:"email_verified,omitempty" toml:"email_verified" yaml:"email_verified,omitempty"`
	Mobile         null.String `boil:"mobile" json:"mobile,omitempty" toml:"mobile" yaml:"mobile,omitempty"`
	MobileVerified null.Bool   `boil:"mobile_verified" json:"mobile_verified,omitempty" toml:"mobile_verified" yaml:"mobile_verified,omitempty"`
}

type AccountSlice []Account

type LoginRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	IsOTP    bool   `json:"is_otp"`
}

type LoginResponse struct {
	Token []byte `json:"token"`
}

type VerifyEmailRequest struct {
	Email string `json:"email"`
}

type CheckEmailVerification struct {
	VerifyCode string `json:"verify_code"`
}

type VerifyMobileRequest struct {
	Mobile string `json:"mobile"`
}

type CheckMobileVerification struct {
	VerifyCode string `json:"verify_code"`
}
