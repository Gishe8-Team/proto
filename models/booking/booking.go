package booking

// this is the message that events sends when cron job runs and booking is started
type InitBooking struct {
	Meta
}

type BookingEvent int

const (
	Start BookingEvent = iota + 1
	Pause
	Stop
	Finalize
)

// all the changes on status of timeslots from events are using this struct to change the booking status
type ChangeBookingState struct {
	TimeslotID string       `json:"timeslot_id"`
	State      BookingEvent `json:"state"`
}
