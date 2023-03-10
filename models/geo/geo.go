package models

type RequestState struct {
	StateID   int    `json:"state_id,omitempty"`
	StateName string `json:"state_name,omitempty"`
}

type RequestCity struct {
	StateID  int    `json:"state_id,omitempty"`
	CityID   int    `json:"city_id,omitempty"`
	CityName string `json:"city_name,omitempty"`
}

type ResponseState struct {
	ConnectionID string `json:"connection_id"`
	ResponseCode int8   `json:"response_code"`
	Message      string `json:"message"`

	States []StateModel `json:"states,omitempty"`
}

type ResponseCity struct {
	ConnectionID string `json:"connection_id"`
	ResponseCode int8   `json:"response_code"`
	Message      string `json:"message"`

	Cities []CityModel  `json:"cities,omitempty"`
	States []StateModel `json:"states,omitempty"`
}

type CityModel struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	StateID int    `json:"state_id,omitempty"`
}

type StateModel struct {
	ID     int           `json:"id"`
	Name   string        `json:"name"`
	Cities *[]*CityModel `json:"cities,omitempty"`
}
