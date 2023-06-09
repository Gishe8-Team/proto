package pay

import (
	"github.com/Gishe8-Team/proto/models/accounting"
	"github.com/Gishe8-Team/proto/models/marketing"
	"github.com/Gishe8-Team/proto/models/payement"
)

type Pay struct {
	marketing.Discount `json:"discount"`
	accounting.Faktor  `json:"faktor"`
	Checkout           payement.Checkouts `json:"checkout"`
}
