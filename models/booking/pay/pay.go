package pay

import (
	"github.com/Gishe8-Team/proto/models/accounting"
	"github.com/Gishe8-Team/proto/models/payement"
)

type Pay struct {
	Faktor   accounting.Faktor  `json:"faktor"`
	Checkout payement.Checkouts `json:"checkout"`
}
