package complex

type ArrangementModel struct {
	ArrangementID   string `json:"arrangement_id"`
	Name            string `json:"name"`
	Creator         string `json:"creator"`
	HallID          string `json:"hall_id"`
	IsCustom        bool   `json:"is_custom"`
	SeatCount       int    `json:"seat_count"`
	HaveBalcony     bool   `json:"have_balcony"`
	SeatArrangement []byte `json:"seat_arrangement"`
}

type SeatFormationModel struct {
	Zones []Zone `json:"zones"`
}

type Zone struct {
	ZoneID    string `json:"zone_id"`
	Available bool   `json:"available"`
	Rows      []Row  `json:"rows"`
}

type Row struct {
	RowNumber int    `json:"row_number"`
	Available bool   `json:"available"`
	Seats     []Seat `json:"seats"`
}

// Seat struct define seats metadata
type Seat struct {
	//SeatNumber show number of seat in row
	SeatNumber int64 `json:"seat_number"`
	//SeatId show seat unique id
	SeatID string `json:"seat_id"`
	//Status define seat status [available, reserved, booked]
	Status string `json:"status"`
	//PriceGroup show this seat is member of which pricing group
	PriceGroup string `json:"price_group"`
	//ProfileID show which user booked this seat
	ProfileID string `json:"profile_id"`
	//SeatType show which seat type [vip, first class, second class, etc]
	SeatType string `json:"seat_type"`
}