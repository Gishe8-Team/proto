package event

type Category struct {
	ID     string `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name   string `boil:"name" json:"name" toml:"name" yaml:"name"`
	Title  string `boil:"title" json:"title" toml:"title" yaml:"title"`
	Parent string `boil:"parent" json:"parent,omitempty" toml:"parent" yaml:"parent,omitempty"`
}

type QueryCategory struct {
	Name   string `json:"name"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}
