package accounting

// FaktorTypeModel is an object representing the database table.
type FaktorTypeModel struct {
	TypeID          int16  `boil:"type_id" json:"type_id" toml:"type_id" yaml:"type_id"`
	TypeName        string `boil:"type_name" json:"type_name" toml:"type_name" yaml:"type_name"`
	TypeDescription string `boil:"type_description" json:"type_description" toml:"type_description" yaml:"type_description"`
}
