package accounting

import "time"

type FindFaktorRequest struct {
	UserID         string    `json:"user_id,omitempty"`
	EventID        string    `json:"event_id,omitempty"`
	StartTime      time.Time `json:"start_time,omitempty"`
	EndTime        time.Time `json:"end_time,omitempty"`
	NotPayedFaktor bool      `json:"not_payed_faktor,omitempty"`
	Offset         int       `json:"offset"`
	Limit          int       `json:"limit"`
}

type FindFaktorResponse struct {
	Faktors []*Faktor `json:"faktors"`
	Status  int       `json:"status"`
	//ItemCount  int       `json:"item_count"`
	//PageNumber int       `json:"page_number"`
	//PageSize   int       `json:"page_size"`
}
