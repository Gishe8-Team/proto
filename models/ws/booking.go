package ws

type BookSeat struct {
	TimeslotID string `json:"timeslot_id"`
	Hash       string `json:"hash"`
	SeatID     string `json:"seat_id"`
}

type SeatBooked struct {
	TimeslotID string `json:"timeslot_id"`
	Hash       string `json:"hash"`
	SeatID     string `json:"seat_id"`
}
