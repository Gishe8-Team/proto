package booking

import (
	"github.com/Gishe8-Team/proto/models/event"
	"time"
)

type Meta struct {
	TimeSlotID  string             `json:"time_slot_id"`
	StartTime   time.Time          `json:"start_time"`
	EndTime     time.Time          `json:"end_time"`
	IsActive    bool               `json:"is_active"`
	PriceGroups []event.PriceGroup `json:"price_groups"`
	SeatTypes   []event.SeatType   `json:"seat_types"`
}
