package complex

type ComplexModel struct {
	ComplexID   string   `boil:"id" json:"complex_id"`
	Name        string   `boil:"name" json:"name"`
	Geolocation Geoloc   `boil:"geolocation" json:"geolocation"`
	Address     string   `boil:"address" json:"address"`
	Email       string   `boil:"email" json:"email"`
	Phone       string   `boil:"phone" json:"phone"`
	Description string   `boil:"description" json:"description"`
	ImagesURL   []string `boil:"images_url,bind" json:"images_url"`
	CityID      int      `boil:"city_id" json:"city_id"`
	StateID     int      `boil:"ostan_id" json:"state_id"`
}

type Geoloc struct {
	Latitude  float64 `boil:"latitude" json:"latitude"`
	Longitude float64 `boil:"longitude" json:"longitude"`
}
