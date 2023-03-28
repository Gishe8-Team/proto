package sso

import "github.com/volatiletech/null/v8"

type AccessRoute struct {
	ID   int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
}

type AccessRouteSlice []AccessRoute
