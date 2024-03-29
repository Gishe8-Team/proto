package event

import "github.com/volatiletech/null/v8"

type CategoryModel struct {
	ID     string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name   string      `boil:"name" json:"name" toml:"name" yaml:"name"`
	Title  string      `boil:"title" json:"title" toml:"title" yaml:"title"`
	Image  null.String `boil:"image" json:"image"`
	Parent null.String `boil:"parent" json:"parent,omitempty" toml:"parent" yaml:"parent,omitempty"`
	TypeID null.String `boil:"type_id" json:"type_id,omitempty" toml:"type_id" yaml:"type_id,omitempty"`
}

type QueryCategoryModel struct {
	ID     null.String `json:"id"`
	Name   null.String `json:"name"`
	Limit  null.Int    `json:"limit"`
	Offset null.Int    `json:"offset"`
}
