package dms

import "github.com/volatiletech/null/v8"

type Document struct {
	ID        string `boil:"id" json:"id" toml:"id" yaml:"id"`
	OwnerID   string `boil:"owner_id" json:"owner_id" toml:"owner_id" yaml:"owner_id"`
	URL       string `json:"url" boil:"location"`
	Type      string `boil:"type" json:"dms_type" toml:"dms_type" yaml:"dms_type"`
	CreatedAt int64  `boil:"created_at" json:"created_at,omitempty"`
}

type DocumentQuery struct {
	ID              null.String `boil:"id" json:"id" toml:"id" yaml:"id"`
	OwnerID         null.String `boil:"owner_id" json:"owner_id" toml:"owner_id" yaml:"owner_id"`
	Type            null.String `boil:"type" json:"dms_type" toml:"dms_type" yaml:"dms_type"`
	CreatedAt       null.Int64  `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	CreatedByAtTill null.Int64  `json:"created_by_at_till"`
	Limit           null.Int    `json:"limit"`
	Offset          null.Int    `json:"offset"`
}

type DocType struct {
	ID    string `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name  string `boil:"name" json:"name" toml:"name" yaml:"name"`
	Title string `boil:"title" json:"title" toml:"title" yaml:"title"`
}

type DocTypeQuery struct {
	Limit  null.Int `json:"limit"`
	Offset null.Int `json:"offset"`
}
