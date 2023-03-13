package event

type Score struct {
	EventID string  `json:"event_id"`
	Score   float64 `json:"score"`
}

type QueryScore struct {
	EventID  string  `json:"eventId"`
	UserID   string  `json:"userId"`
	MinScore float64 `json:"minScore"`
	MaxScore float64 `json:"maxScore"`
	Limit    int     `json:"limit"`
	Offset   int     `json:"offset"`
}
