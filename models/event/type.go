package event

import "github.com/volatiletech/null/v8"

type EventsTypeModel struct {
	ID    string      `boil:"type_id" json:"id" toml:"id" yaml:"id"`
	Name  string      `boil:"type_name" json:"name" toml:"name" yaml:"name"`
	Image null.String `boil:"image" json:"image"`
	Title string      `boil:"type_title" json:"title" toml:"title" yaml:"title"`
}

type QueryEventsTypeModel struct {
	Limit  null.Int `json:"limit"`
	Offset null.Int `json:"offset"`
}
