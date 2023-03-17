package profile

type GenderType struct {
	ID   int16  `boil:"id" json:"id,omitempty"`
	Name string `boil:"name" json:"name"`
}
