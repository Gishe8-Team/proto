package event

import "time"

type TimeSlot struct {
	ID            string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	EventID       string    `boil:"event_id" json:"event_id" toml:"event_id" yaml:"event_id"`
	HallID        string    `boil:"hall_id" json:"hall_id" toml:"hall_id" yaml:"hall_id"`
	ArrangementID string    `boil:"arrangement_id" json:"arrangement_id" toml:"arrangement_id" yaml:"arrangement_id"`
	StartTime     time.Time `boil:"start_time" json:"start_time" toml:"start_time" yaml:"start_time"`
	EndTime       time.Time `boil:"end_time" json:"end_time" toml:"end_time" yaml:"end_time"`
	Tnos          int       `boil:"tnos" json:"tnos" toml:"tnos" yaml:"tnos"`
	Fnos          int       `boil:"fnos" json:"fnos" toml:"fnos" yaml:"fnos"`
	Public        bool      `boil:"public" json:"public" toml:"public" yaml:"public"`
	CreatedAt     time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
}

type QueryTimeSlot struct {
	EventID   string `json:"event_id"`
	HallID    string `json:"hall_id"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
	Limit     int    `json:"limit"`
	Offset    int    `json:"offset"`
}
