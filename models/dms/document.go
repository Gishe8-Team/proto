package dms

type QueryDocumentsModels struct {
	ID      string `boil:"id" json:"id"`
	OwnerID string `boil:"owner_id" json:"owner_id"`
	DMSType string `boil:"dms_type" json:"type"`
}
