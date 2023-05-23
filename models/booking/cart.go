package booking

import "github.com/Gishe8-Team/proto/models/event"

// RequestViewCartModel this request in step 1 of booking
type RequestViewCartModel struct {
	TimeslotID string `json:"timeslot_id"`
	UserID     string `json:"user_id"`
}

// ResponseViewCartModelLayout this is the response to step 1 of booking if booking have layout
type ResponseViewCartModelLayout struct {
	Cart     *CartModel   `json:"cart"`
	OldCarts []CartModel  `json:"old_cart,omitempty"`
	Meta                  // Status of cart {open, pending, paid}
	Seats    []event.Seat `json:"seats"`
}

// ResponseViewCartModelNoLayout this is the response to step 1 of booking if booking haven't layout
type ResponseViewCartModelNoLayout struct {
	Cart *CartModel `json:"cart"`
	Meta
	Zones []ZoneModel `json:"zones"`
}

type ResponseViewCartModelRedirectFactor struct {
	Cart *CartModel `json:"cart"`
}

// ZoneModel composite to above struct
type ZoneModel struct {
	ZoneID int      `json:"layout_id"`
	Prices []Prices `json:"prices"`
}

// Prices composite to above struct
type Prices struct {
	event.PriceGroup
	FreeSeats int `json:"free_seats"`
}

// RequestToggleSeat income from WebSocket Gateway and request do a seat of desired timeslot
// reserve/unreserved  to/from a cart
type RequestToggleSeat struct {
	TimeslotID string `json:"timeslot_id"` // The ID of the timeslot selected by the customer.
	Hash       string `json:"hash"`
	SeatID     string `json:"seat_id"`
}

// ResponseToggleSeat publish to all users subscribed in timeslot channel
type ResponseToggleSeat struct {
	TimeslotID string `json:"timeslot_id"` // The ID of the timeslot selected by the customer.
	Hash       string `json:"hash"`
	SeatID     string `json:"seat_id"`
}

// RequestSubmitCartModel represents the data required to finalize a customer's cart.
type RequestSubmitCartModel struct {
	TimeslotID string          `json:"timeslot_id"` // The ID of the timeslot selected by the customer.
	UserID     string          `json:"user_id"`     // The ID of the user who owns the cart.
	Hash       string          `json:"hash"`
	Coupon     string          `json:"coupon"`          // An optional coupon code to apply to the cart.
	Seats      []NoLayoutSeats `json:"seats,omitempty"` // if seats exist then add seats per price group (no layout selection)
}

// NoLayoutSeats composite in above struct
type NoLayoutSeats struct {
	ZoneID       int `json:"zone_id"`
	PriceGroupID int `json:"price_group_id"`
	Count        int `json:"count"`
}

// CartModel define model of cart
type CartModel struct {
	UserID     string       `json:"user_id,omitempty"`
	TimeSlotID string       `json:"time_slot_id,omitempty"`
	Hash       string       `json:"hash,omitempty"`
	IdleTime   int64        `json:"idle_time,omitempty"` // TTL Cart it should create on returning timeslot meta
	Status     CartStatus   `json:"status"`              // Status of cart {open, pending, paid}
	Seats      []event.Seat `json:"seats"`
}

// ChangeCartStatus for changing the status of a cart for reflecting
// operation of cart like pending for payment or successful payment.
type ChangeCartStatus struct {
	UserID     string     `json:"user_id,omitempty"`
	TimeSlotID string     `json:"time_slot_id,omitempty"`
	Hash       string     `json:"hash,omitempty"`
	State      CartStatus `json:"status"`
}
