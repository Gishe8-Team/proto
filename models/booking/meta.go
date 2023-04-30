package booking

import (
	complex2 "github.com/Gishe8-Team/proto/models/complex"
	"github.com/Gishe8-Team/proto/models/event"
	"time"
)

type Meta struct {
	TimeSlotID  string                    `json:"time_slot_id"`
	Name        string                    `json:"name"`
	Description string                    `json:"description"`
	StartTime   time.Time                 `json:"start_time"`
	EndTime     time.Time                 `json:"end_time"`
	Event       event.ViewSmallEventModel `json:"event"`
	Hall        complex2.HallModel        `json:"hall"`
	IsActive    bool                      `json:"is_active"`
	Background  string                    `json:"background"`
	PriceGroups []event.PriceGroup        `json:"price_groups"`
	SeatTypes   []event.SeatType          `json:"seat_types"`
	Layouts     []event.Layout            `json:"layouts"`
}
