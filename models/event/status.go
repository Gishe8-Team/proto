package event

import "github.com/volatiletech/null/v8"

type EventsStatusModel struct {
	ID    string `boil:"status_id" json:"id" toml:"id" yaml:"id"`
	Name  string `boil:"status_name" json:"name" toml:"name" yaml:"name"`
	Title string `boil:"status_title" json:"title" toml:"title" yaml:"title"`
}

type QueryEventsStatusModel struct {
	Name   null.String `json:"name"`
	Limit  null.Int    `json:"limit"`
	Offset null.Int    `json:"offset"`
}
