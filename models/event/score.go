package event

import "github.com/volatiletech/null/v8"

type EventsScoreModel struct {
	EventID string  `json:"event_id"`
	Score   float64 `json:"score"`
}

type QueryEventsScoreModel struct {
	EventID  null.String  `json:"eventId"`
	UserID   null.String  `json:"userId"`
	MinScore null.Float64 `json:"minScore"`
	MaxScore null.Float64 `json:"maxScore"`
	Limit    null.Int     `json:"limit"`
	Offset   null.Int     `json:"offset"`
}
