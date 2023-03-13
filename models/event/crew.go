package event

type Crew struct {
	ID     string   `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name   string   `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	Bio    string   `boil:"bio" json:"bio,omitempty" toml:"bio" yaml:"bio,omitempty"`
	Images []string `boil:"images" json:"images" toml:"images" yaml:"images"`
}

type QueryCrew struct {
	EventID string `json:"event_id"`
	Name    string `json:"name"`
	Limit   int    `json:"limit"`
	Offset  int    `json:"offset"`
}
