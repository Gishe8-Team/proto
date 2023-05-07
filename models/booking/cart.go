package booking

import "github.com/Gishe8-Team/proto/models/event"

type RequestViewCartModel struct {
	TimeslotID string `json:"timeslot_id"`
	UserID     string `json:"user_id"`
}

// ResponseViewCartModelLayout this is the response to state one of booking
type ResponseViewCartModelLayout struct {
	Cart    Cart         `json:"cart"`
	OldCart Cart         `json:"old_cart,omitempty"`
	Meta                 // Status of cart {open, pending, paid}
	Seats   []event.Seat `json:"seats"`
}

type ResponseViewCartModelNoLayout struct {
	Cart Cart `json:"cart"`
	Meta
	Zones []struct {
		ZoneID int      `json:"zone_id"`
		Prices []Prices `json:"prices"`
	} `json:"zones"`
}

type RequestBookingNoLayout struct {
	Hash                      string `json:"hash"`	      // The Hash of the user cart.
	Seats []struct {
		ZoneID       	int `json:"zone_id"`
		PriceGroupID 	int `json:"price_group_id"`
		Count        	int `json:"count"`
	} `json:"seats"`
}
	
type Prices struct {
	event.PriceGroup
	FreeSeats int `json:"free-seats"`
}

type FinalizeCartModelNoLayout struct {
	Seats []struct {
		ZoneID       int `json:"zone_id"`
		PriceGroupID int `json:"price_group_id"`
		Count        int `json:"count"`
	} `json:"seats,omitempty"`
}

// FinalizeCartModel represents the data required to finalize a customer's cart.
type RequestFinalizeCartModel struct {
	TimeslotID                string `json:"timeslot_id"` // The ID of the timeslot selected by the customer.
	UserID                    string `json:"user_id"`     // The ID of the user who owns the cart.
	Hash                      string `json:"hash"`
	Coupon                    string `json:"coupon"` // An optional coupon code to apply to the cart.
	FinalizeCartModelNoLayout        // if seats exist then add seats per price group (no layout selection)
}

type Cart struct {
	UserID     string       `json:"user_id,omitempty"`
	TimeSlotID string       `json:"time_slot_id,omitempty"`
	Hash       string       `json:"hash,omitempty"`
	IdleTime   int64        `json:"idle_time,omitempty"` // TTL Cart it should create on returning timeslot meta
	Status     CartStatus   `json:"status"`              // Status of cart {open, pending, paid}
	Seats      []event.Seat `json:"seats"`
}
