package accounting

// FaktorStatusModel is an object representing the database table.
type FaktorStatusModel struct {
	StatusID   int16  `boil:"status_id" json:"status_id" toml:"status_id" yaml:"status_id"`
	StatusName string `boil:"status_name" json:"status_name" toml:"status_name" yaml:"status_name"`
}
