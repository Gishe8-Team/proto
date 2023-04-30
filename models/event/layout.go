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
	X          int     `json:"x"`
	Y          int     `json:"y"`
	ZoneID     int8    `json:"zone_id"`
	Row        int8    `json:"row"`
	Status     TStatus `json:"status"`
	PriceGroup int8    `json:"price_group"`
	SeatType   int8    `json:"seat_type"`
	UserID     string  `json:"user_id"`
}

type PriceGroup struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type SeatType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Layout struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	X           int    `json:"x"`
	Y           int    `json:"y"`
	Height      int    `json:"height"`
	Width       int    `json:"width"`
	Types       string `json:"types"`
	BorderColor string `json:"border_color"`
	FillColor   string `json:"fill_color"`
}

type TimeSlotLayout struct {
	ID         string       `json:"id"`
	PriceGroup []PriceGroup `json:"price_group"`
	SeatType   []SeatType   `json:"seat_type"`
	Seats      []Seat       `json:"seats"`
}
