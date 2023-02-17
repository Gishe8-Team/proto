package models

type ArrangementModel struct {
	ArrangementID   string `json:"arrangement_id"`
	Name            string `json:"name"`
	Creator         string `json:"creator"`
	HallID          string `json:"hall_id"`
	SeatCount       int    `json:"seat_count"`
	HaveBalcony     bool   `json:"have_balcony"`
	SeatArrangement []byte `json:"seat_arrangement"`
}

type SeatFormationModel struct {
}
