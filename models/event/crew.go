package event

import (
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/types"
)

type EventCrewModel struct {
	CrewModel
	Role null.String `boil:"role" json:"role,omitempty"`
}

type CrewModel struct {
	ID     string            `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name   null.String       `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	Bio    null.String       `boil:"bio" json:"bio,omitempty" toml:"bio" yaml:"bio,omitempty"`
	Images types.StringArray `boil:"images" json:"images" toml:"images" yaml:"images"`
}

type QueryCrewModel struct {
	ID      string      `json:"id"`
	EventID null.String `json:"event_id"`
	Name    null.String `json:"name"`
	Limit   null.Int    `json:"limit"`
	Offset  null.Int    `json:"offset"`
}
