package event

import (
	"github.com/Gishe8-Team/proto/models/complex"
	"github.com/Gishe8-Team/proto/models/geo"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/types"
)

type ViewFullEventModel struct {
	ID          string              `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name        string              `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	Description null.String         `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`
	Score       float64             `boil:"score" json:"score,omitempty"`
	Type        EventsTypeModel     `boil:"type,bind" json:"type" toml:"type" yaml:"type"`
	Status      EventsStatusModel   `boil:"status,bind." json:"status,omitempty" toml:"status" yaml:"status"`
	Hall        []complex.HallModel `boil:",bind" json:"hall" toml:"hall_id" yaml:"hall_id"`
	State       []geo.StateModel    `boil:",bind" json:"state" toml:"state" yaml:"state"`
	City        []geo.CityModel     `boil:",bind" json:"city" toml:"city" yaml:"city"`
	Category    []CategoryModel     `boil:",bind" json:"category"`
	Crew        []EventCrewModel    `boil:",bind" json:"crew"`
	TimeSlots   []TimeslotModel     `boil:",bind" json:"timeslots"`
	Cover       string              `boil:"cover" json:"cover" toml:"cover" yaml:"cover"`
	Poster      string              `boil:"poster" json:"poster" toml:"poster"`
	Gallery     []string            `boil:"images" json:"images" toml:"images" yaml:"images"`
}

type ViewSmallEventModel struct {
	ID          string            `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name        string            `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	Description null.String       `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`
	Score       float64           `boil:"score" json:"score,omitempty"`
	Type        EventsTypeModel   `boil:"type,bind" json:"type" toml:"type" yaml:"type"`
	Status      EventsStatusModel `boil:"status,bind" json:"status,omitempty" toml:"status" yaml:"status"`
	Cover       string            `boil:"cover" json:"cover" toml:"cover" yaml:"cover"`
	Poster      string            `boil:"poster" json:"poster" toml:"poster"`
}

type EventModel struct {
	ID          string            `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name        string            `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	Description null.String       `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`
	Type        string            `boil:"type_id" json:"type_id" toml:"type" yaml:"type"`
	Status      string            `boil:"status_id" json:"status_id,omitempty" toml:"status" yaml:"status"`
	HallID      types.StringArray `boil:"hall_id" json:"hall_id" toml:"hall_id" yaml:"hall_id"`
	Wage        float64           `boil:"wage" json:"wage" `
	State       types.StringArray `json:"state_id" toml:"state" yaml:"state"`
	City        types.StringArray `json:"city_id" toml:"city" yaml:"city"`
	Cover       string            `boil:"cover" json:"cover" toml:"cover" yaml:"cover"`
	Poster      string            `boil:"poster" json:"poster" toml:"poster"`
	Gallery     []string          `boil:"images" json:"images" toml:"images" yaml:"images"`
}

type QueryEventModel struct {
	ID       null.String       `json:"id"`
	Name     null.String       `json:"name,omitempty"`
	Type     null.String       `json:"type_id"`
	Status   null.String       `json:"status_id,omitempty"`
	Hall     types.StringArray `json:"hall_id"`
	State    types.StringArray `json:"state_id"`
	City     types.StringArray `json:"city_id"`
	Category types.StringArray `json:"category_id"`
	Crew     types.StringArray `json:"crew_id"`
	MinScore null.Float64      `json:"minScore"`
	MaxScore null.Float64      `json:"maxScore"`
	Limit    null.Int          `json:"limit"`
	Offset   null.Int          `json:"offset"`
}

type ViewLanding struct {
	Body []EventViewWidget `json:"body"`
}

type EventViewWidget struct {
	Type  string       `json:"type"`
	Title string       `json:"title"`
	Data  []EventModel `json:"data"`
}
