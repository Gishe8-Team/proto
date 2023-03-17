package dms

import "github.com/volatiletech/null/v8"

type QueryDocumentsModels struct {
	ID      null.String `boil:"id" json:"id"`
	OwnerID null.String `boil:"owner_id" json:"owner_id"`
	DMSType null.String `boil:"dms_type" json:"type"`
}
