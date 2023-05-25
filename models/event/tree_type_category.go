package event

type EventsTypeCategoryModel struct {
	ID         string              `boil:"type_id" json:"id" toml:"id" yaml:"id"`
	Name       string              `boil:"type_name" json:"name" toml:"name" yaml:"name"`
	Title      string              `boil:"type_title" json:"title" toml:"title" yaml:"title"`
	Categories []*CategorySumModel `boil:"-" json:"categories,omitempty" toml:"-" yaml:"-"`
}

type CategorySumModel struct {
	ID    string `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name  string `boil:"name" json:"name" toml:"name" yaml:"name"`
	Title string `boil:"title" json:"title" toml:"title" yaml:"title"`
}
