package complex

type HallModel struct {
	HallID      string   `json:"hall_id"`
	Name        string   `json:"name"`
	ComplexID   string   `json:"complex_id"`
	Images      []string `json:"images"`
	SeatCount   int      `json:"seat_count"`
	HaveBalcony bool     `json:"have_balcony"`
}
