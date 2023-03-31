package sso

import "github.com/volatiletech/null/v8"

type LoginArchive struct {
	ID        int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	IPAddress null.String `boil:"ip_address" json:"ip_address,omitempty" toml:"ip_address" yaml:"ip_address,omitempty"`
	Time      null.Time   `boil:"time" json:"time,omitempty" toml:"time" yaml:"time,omitempty"`
	UserID    string      `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	Device    null.String `boil:"device" json:"device,omitempty" toml:"device" yaml:"device,omitempty"`
}

type LoginArchiveSlice *[]*LoginArchive
