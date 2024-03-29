package booking

import "github.com/Gishe8-Team/proto/models/event"

// CreateBookingSlot this is the message that events sends when cron job runs and booking is started
type CreateBookingSlot struct {
	Meta
	HallID    string       `json:"hall_id"`
	ComplexID string       `json:"complex_id"`
	Seats     []event.Seat `json:"seats,omitempty"`
}

// ChangeBookingSlotStatus all the changes on status of timeslots from events are using this struct to change the booking status
type ChangeBookingSlotStatus struct {
	TimeslotID string               `json:"timeslot_id"`
	Status     event.TimeSlotStatus `json:"status"`
}

// GetBookingSlot for requesting a single timeslot booking data
type GetBookingSlot struct {
	TimeslotID string `json:"timeslot_id"`
}

type ResponseSingleBookingSlot struct {
	Meta
	Seats []event.Seat `json:"seats,omitempty"`
	//Carts []Cart       `json:"carts"` // is it possible to have this and can we have a service for CRUD on booking carts?
}
