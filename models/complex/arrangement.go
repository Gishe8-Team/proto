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
}
