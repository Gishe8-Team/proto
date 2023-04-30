package booking

import "github.com/Gishe8-Team/proto/models/event"

type RequestViewCartModel struct {
	TimeslotID string `json:"timeslot_id"`
	UserID     string `json:"user_id"`
}

// ViewCartModel this is the response to state one of booking
type ResponseViewCartModel struct {
	Cart
	OldCart Cart `json:"oldCart,omitempty"`
	Meta
}

// FinalizeCartModel represents the data required to finalize a customer's cart.
type FinalizeCartModel struct {
	TimeslotID string `json:"timeslot_id"` // The ID of the timeslot selected by the customer.
	UserID     string `json:"user_id"`     // The ID of the user who owns the cart.
	Coupon     string `json:"coupon"`      // An optional coupon code to apply to the cart.
}

type Cart struct {
	UserID     string       `json:"user_id,omitempty"`
	TimeSlotID string       `json:"time_slot_id,omitempty"`
	Hash       string       `json:"hash,omitempty"`
	IdleTime   int64        `json:"idle_time,omitempty"` // TTL Cart it should create on returning timeslot meta
	Status     string       `json:"status"`              // Status of cart {open, pending, paid}
	Seats      []event.Seat `json:"seats"`
}
