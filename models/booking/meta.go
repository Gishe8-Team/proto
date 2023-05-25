package booking

import (
	complex2 "github.com/Gishe8-Team/proto/models/complex"
	"github.com/Gishe8-Team/proto/models/event"
	"time"
)

type Meta struct {
	TimeSlotID  string                    `json:"time_slot_id"`
	Name        string                    `json:"name,omitempty"`
	Description string                    `json:"description,omitempty"`
	StartTime   time.Time                 `json:"start_time,omitempty"`
	EndTime     time.Time                 `json:"end_time,omitempty"`
	Event       event.ViewSmallEventModel `json:"event,omitempty"`
	Hall        complex2.HallModel        `json:"hall,omitempty"`
	Status      event.TimeSlotStatus      `json:"status,omitempty"`
	Background  string                    `json:"background,omitempty"`
	PriceGroups []event.PriceGroup        `json:"price_groups,omitempty"`
	BookingMode event.BookingMode         `json:"booking_mode,omitempty"`
	SeatTypes   []event.SeatType          `json:"seat_types,omitempty"`
	Layouts     []event.Layout            `json:"layouts,omitempty"`
	TermsLimit  string                    `json:"terms_limits"`
}
