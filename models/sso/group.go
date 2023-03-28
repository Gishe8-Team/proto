package sso

import "github.com/volatiletech/null/v8"

type Group struct {
	ID   int16       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	ACL  null.String `boil:"acl" json:"acl,omitempty" toml:"acl" yaml:"acl,omitempty"`
}

type GroupSlice []Group
