package proto

type RequestState struct {
	ConnectionID string `json:"connection_id"`

	StateID   int    `json:"state_id,omitempty"`
	StateName string `json:"state_name,omitempty"`
}

type RequestCity struct {
	ConnectionID string `json:"connection_id"`

	StateID  int    `json:"state_id,omitempty"`
	CityID   int    `json:"city_id,omitempty"`
	CityName string `json:"city_name,omitempty"`
}

const (
	AddState = iota + 521
	EditState
	RemoveState
	FindStateById
	ListState
	ListStatesWithCities
)

const (
	AddCity = iota + 531
	EditCity
	RemoveCity
	FindCityByStateId
	FindCityByCityId
)

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
	ID      string `json:"id"`
	Name    string `json:"name"`
	StateID string `json:"state_id,omitempty"`
}

type StateModel struct {
	ID     string      `json:"id"`
	Name   string      `json:"name"`
	Cities []CityModel `json:"cities,omitempty"`
}
