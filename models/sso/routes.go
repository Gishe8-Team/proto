package sso

type AccessRoute struct {
	ID   int64  `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name string `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
}

type AccessRouteSlice *[]*AccessRoute
