package event

import (
	"github.com/Gishe8-Team/proto/models/complex"
	"github.com/Gishe8-Team/proto/models/geo"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/types"
)

type ViewFullEventModel struct {
	ID          string                 `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name        string                 `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	Description null.String            `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`
	Score       float64                `boil:"score" json:"score,omitempty"`
	Type        EventsTypeModel        `boil:",bind" json:"type" toml:"type" yaml:"type"`
	Status      EventsStatusModel      `boil:",bind." json:"status,omitempty" toml:"status" yaml:"status"`
	Complex     []complex.ComplexModel `boil:"-" json:"complex"`
	HallIDs     types.StringArray      `boil:"hall_id" json:"hall_ids"`
	Hall        []complex.HallModel    `boil:"-" json:"halls"`
	//StateID can be deleted. stateID exists in cityModel
	StateIDs types.Int64Array `boil:"state_id" json:"state_ids"`
	State    []geo.StateModel `boil:"-" json:"states"`
	//cityID must be int array
	CityIDs   types.Int64Array            `boil:"city_id" json:"city_ids"`
	City      []geo.CityModel             `boil:"-" json:"cities"`
	Category  []*CategoryModel            `boil:"-" json:"category"`
	Crew      []*ViewSingleEventCrewModel `boil:"-" json:"crew"`
	TimeSlots []*TimeslotModel            `boil:"-" json:"timeslots"`
	Cover     string                      `boil:"cover" json:"cover" toml:"cover" yaml:"cover"`
	Poster    string                      `boil:"poster" json:"poster" toml:"poster"`
	Gallery   types.StringArray           `boil:"images" json:"images" toml:"images" yaml:"images"`
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
	Fee         float64           `boil:"wage" json:"wage" `
	State       types.Int64Array  `json:"state_id" toml:"state" yaml:"state"`
	City        types.Int64Array  `json:"city_id" toml:"city" yaml:"city"`
	Cover       string            `boil:"cover" json:"cover" toml:"cover" yaml:"cover"`
	Poster      string            `boil:"poster" json:"poster" toml:"poster"`
	Gallery     types.StringArray `boil:"gallery" json:"images" toml:"images" yaml:"images"`
	Score       float64           `boil:"score" json:"rate"`
}

type QueryEventModel struct {
	ID       null.String       `json:"id"`
	Slug     null.String       `json:"slug,omitempty"`
	Name     null.String       `json:"name,omitempty"`
	Type     null.String       `json:"type_id"`
	Status   null.String       `json:"status_id,omitempty"`
	Hall     types.StringArray `json:"hall_id"`
	State    types.Int64Array  `json:"state_id"`
	City     types.Int64Array  `json:"city_id"`
	Category types.StringArray `json:"category_id"`
	Crew     types.StringArray `json:"crew_id"`
	MinScore null.Float64      `json:"min_score"`
	MaxScore null.Float64      `json:"max_score"`
	Limit    null.Int          `json:"limit"`
	Offset   null.Int          `json:"offset"`
}

type ViewLanding struct {
	Body []*EventViewWidget `json:"body"`
}

type EventViewWidget struct {
	Type  string      `json:"type"`
	Title string      `json:"title"`
	Data  interface{} `json:"data"`
}
