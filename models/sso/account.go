package sso

import "github.com/volatiletech/null/v8"

type Account struct {
	ID               string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserName         null.String `boil:"user_name" json:"user_name,omitempty" toml:"user_name" yaml:"user_name,omitempty"`
	Password         null.String `boil:"password" json:"password,omitempty" toml:"password" yaml:"password,omitempty"`
	GroupID          int16       `boil:"group_id" json:"group_id" toml:"group_id" yaml:"group_id"`
	UserACL          string      `boil:"user_acl" json:"user_acl" toml:"user_acl" yaml:"user_acl"`
	UserEmail        null.String `boil:"user_email" json:"user_email,omitempty" toml:"user_email" yaml:"user_email,omitempty"`
	IsEmailVerified  null.Bool   `boil:"is_email_verified" json:"is_email_verified,omitempty" toml:"is_email_verified" yaml:"is_email_verified,omitempty"`
	UserMobile       null.String `boil:"user_mobile" json:"user_mobile,omitempty" toml:"user_mobile" yaml:"user_mobile,omitempty"`
	IsMobileVerified null.Bool   `boil:"is_mobile_verified" json:"is_mobile_verified,omitempty" toml:"is_mobile_verified" yaml:"is_mobile_verified,omitempty"`
	CreatedAt        null.Time   `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt        null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	DeletedAt        null.Time   `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
}

type AccountSlice []Account

type LoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	IsOTP    bool   `json:"is_otp"`
}

type LoginResponse struct {
	Token []byte `json:"token"`
}
