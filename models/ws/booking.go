package ws

type BookingSeatEvent struct {
	TimeslotID string `json:"timeslot_id"`
	Hash       string `json:"hash"`
	SeatID     string `json:"seat_id"`
	Status     string `json:"status"`
}
