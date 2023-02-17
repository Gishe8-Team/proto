package models

type ComplexModel struct {
	ComplexID   string   `json:"complex_id"`
	Name        string   `json:"name"`
	Geolocation Geoloc   `json:"geolocation"`
	Address     string   `json:"address"`
	Email       string   `json:"email"`
	Phone       string   `json:"phone"`
	Description string   `json:"description"`
	Images      []string `json:"images"`
	CityID      int      `json:"city_id"`
	StateID     int      `json:"state_id"`
}

type Geoloc struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
