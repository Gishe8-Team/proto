package booking

import "github.com/Gishe8-Team/proto/models/event"

// this is the response to state one of booking
type ViewCartModel struct {
	Cart
	OldCart Cart `json:"oldCart,omitempty"`
	Meta
}

// this is the message that ws ms sends to booking when a seat is clicked on client
type ToggleCartModel struct {
	TimeslotID string `json:"timeslot_id"`
	Hash       string `json:"hash"`
	SeatID     string `json:"seat_id"`
}

// this is the request that client sends when done with booking
type FinalizeCartModel struct {
	TimeslotID string `json:"timeslot_id"`
}

type Cart struct {
	UserID     string       `json:"user_id"`
	TimeSlotID string       `json:"time_slot_id"`
	Hash       string       `json:"hash"`
	IdleTime   int64        `json:"idle_time"` // TTL Cart it should create on returning timeslot meta
	Seats      []event.Seat `json:"seats"`
}
