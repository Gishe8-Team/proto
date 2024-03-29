package event

import (
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/types"
)

type ViewSingleEventCrewModel struct {
	CrewModel `boil:",bind"`
	Role      null.String `boil:"role" json:"role,omitempty"`
}

type ViewFullCrewModel struct {
	CrewModel `boil:",bind"`
	Roles     []*ViewCrewRoleModel `boil:"-" json:"roles,omitempty"`
}

type ViewCrewRoleModel struct {
	Role                string `boil:"role" json:"role,omitempty"`
	ViewSmallEventModel `boil:",bind" json:"event"`
}

type CrewRoleModel struct {
	ID       string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	EventID  string      `boil:"event_id" json:"event_id" toml:"event_id" yaml:"event_id"`
	CrewID   string      `boil:"crew_id" json:"crew_id" toml:"crew_id" yaml:"crew_id"`
	Role     null.String `boil:"role" json:"role,omitempty" toml:"role" yaml:"role,omitempty"`
	Position int         `boil:"position" json:"position" toml:"position" yaml:"position"`
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
