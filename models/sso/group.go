package sso

type Group struct {
	ID   int16  `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name string `boil:"name" json:"name" toml:"name" yaml:"name"`
	ACL  string `boil:"acl" json:"acl" toml:"acl" yaml:"acl"`
}

type GroupSlice *[]*Group
