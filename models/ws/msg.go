package ws

// WsBroadcast used for sending message to channels
type Broadcast struct {
	ChannelID string `json:"channel_id"`

	Data []byte `json:"data"`
}

// this used to sned ws events to rmq
type WsEvent struct {
	UserID    string
	ChannelID string
	Event     string `json:"event"`
	Data      []byte
}

// BookSeatMsg request message for booking a seat for particular timeslot.
type BookSeatMsg struct {
	TimeslotID string `json:"timeslot_id"` // ID of the timeslot being booked
	Hash       string `json:"hash"`        // Used to verify authenticity of message
	SeatID     string `json:"seat_id"`     // ID of the seat being booked
}

// BookingSeatEvent represents an event related to booking a seat for a particular timeslot.
type BookingSeatEvent struct {
	TimeslotID string `json:"timeslot_id"` // The ID of the timeslot for which the seat was booked.
	Hash       string `json:"hash"`        // An authentication or authorization token associated with the booking.
	SeatID     string `json:"seat_id"`     // The ID of the booked seat.
	Status     string `json:"status"`      // The current status of the booking (e.g. "confirmed", "pending", "cancelled", etc.).
}
