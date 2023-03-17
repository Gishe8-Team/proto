package event

import (
	"github.com/volatiletech/null/v8"
	"time"
)

type TimeslotModel struct {
	ID            string    `boil:"ts_id" json:"id" toml:"id" yaml:"id"`
	EventID       string    `boil:"ts_event_id" json:"event_id" toml:"event_id" yaml:"event_id"`
	HallID        string    `boil:"ts_hall_id" json:"hall_id" toml:"hall_id" yaml:"hall_id"`
	ArrangementID string    `boil:"ts_arrangement_id" json:"arrangement_id" toml:"arrangement_id" yaml:"arrangement_id"`
	StartTime     time.Time `boil:"ts_start_time" json:"start_time" toml:"start_time" yaml:"start_time"`
	EndTime       time.Time `boil:"ts_end_time" json:"end_time" toml:"end_time" yaml:"end_time"`
	Tnos          int       `boil:"ts_tnos" json:"tnos" toml:"tnos" yaml:"tnos"`
	Fnos          int       `boil:"ts_fnos" json:"fnos" toml:"fnos" yaml:"fnos"`
	Public        bool      `boil:"ts_public" json:"public" toml:"public" yaml:"public"`
	CreatedAt     time.Time `boil:"ts_created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
}

type QueryTimeslotModel struct {
	EventID   null.String `json:"event_id"`
	HallID    null.String `json:"hall_id"`
	StartTime null.Int64  `json:"start_time"`
	EndTime   null.Int64  `json:"end_time"`
	Limit     null.Int    `json:"limit"`
	Offset    null.Int    `json:"offset"`
}
