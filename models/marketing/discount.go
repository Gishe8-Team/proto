package marketing

type Discount struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Percent     float64 `json:"percent"`
	Ceil        float64 `json:"ceil"`
}
