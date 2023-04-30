package booking

import "github.com/Gishe8-Team/proto/models/event"

// ViewCartModel this is the response to state one of booking
type ViewCartModel struct {
	Cart
	OldCart Cart `json:"oldCart,omitempty"`
	Meta
}

// ToggleCartModel this is the message that ws ms sends to booking when a seat is clicked on client
type ToggleCartModel struct {
	TimeslotID string `json:"timeslot_id"`
	Hash       string `json:"hash"`
	SeatID     string `json:"seat_id"`
}

// FinalizeCartModel this is the request that client sends when done with booking
type FinalizeCartModel struct {
	TimeslotID string `json:"timeslot_id"`
	UserID     string `json:"user_id"`
	Coupon     string `json:"coupon"`
}

type Cart struct {
	UserID     string       `json:"user_id,omitempty"`
	TimeSlotID string       `json:"time_slot_id,omitempty"`
	Hash       string       `json:"hash,omitempty"`
	IdleTime   int64        `json:"idle_time,omitempty"` // TTL Cart it should create on returning timeslot meta
	Status     string       `json:"status"`              // Status of cart {open, pending, paid}
	Seats      []event.Seat `json:"seats"`
}
