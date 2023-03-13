package event

type EventsStatus struct {
	ID    string `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name  string `boil:"name" json:"name" toml:"name" yaml:"name"`
	Title string `boil:"title" json:"title" toml:"title" yaml:"title"`
}

type QueryEventsStatus struct {
	Name   string `json:"name"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}
