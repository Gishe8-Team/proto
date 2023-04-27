package booking

import "github.com/Gishe8-Team/proto/models/event"

type Cart struct {
	UserID     string `json:"user_id"`
	TimeSlotID string `json:"time_slot_id"`
	Seats      []event.Seat
}
