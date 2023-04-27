package event

type TStatus int

const (
	Free TStatus = iota + 1
	Blocked
	Reserved
	Booked
)

type Seat struct {
	SeatID     string  `json:"seat_id"`
	SeatNumber int     `json:"seat_number"`
	Zone       string  `json:"zone"`
	Row        int     `json:"row"`
	Status     TStatus `json:"status"`
	PriceGroup string  `json:"price_group"`
	SeatType   string  `json:"seat_type"`
	UserID     string  `json:"user_id"`
}

type PriceGroup struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type SeatType struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type TimeSlotLayout struct {
	ID         string       `json:"id"`
	PriceGroup []PriceGroup `json:"price_group"`
	SeatType   []SeatType   `json:"seat_type"`
	Seats      []Seat       `json:"seats"`
}