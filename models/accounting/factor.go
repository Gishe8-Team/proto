package accounting

import (
	"github.com/volatiletech/null/v8"
)

type FaktorStatus int16

const (
	FAKTOR_PENDING FaktorStatus = iota + 1
	FAKTOR_PAYED
	FAKTOR_CANCELLED
)

//type Factor struct {
//	//Factor Unique ID
//	ID string `json:"id"`
//	//issuing date of factor
//	IssueDate time.Time `json:"issue_date"`
//	//userID
//	IssuingFor string `json:"issuing_for"`
//	//payment status of factor {enum: pending, paid, expired}
//	PaymentStatus string `json:"payment_status"`
//	//sum of taxaddedprice of descriptions
//	TotalPrices float64 `json:"total_prices"`
//	//summary of descriptions discount
//	TotalDiscount float64 `json:"total_discount"`
//	//total amount of factor for paying (= total price - total discount)
//	PayableAmount float64 `json:"payable_amount"`
//	//DOPS is json string contain invoiceDescriptionsSlice
//	DOPS string `json:"dops"`
//	//description of factor
//	Description string `json:"description"`
//	// this should contain userID, CartID ,EventID and more data that needed
//	Data map[string]interface{} `json:"data"`
//}

// Faktor is an object representing the database table.
type Faktor struct {
	ID            string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID        string      `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	CartID        null.String `boil:"cart_id" json:"cart_id,omitempty" toml:"cart_id" yaml:"cart_id,omitempty"`
	StatusID      int16       `boil:"status_id" json:"status_id" toml:"status_id" yaml:"status_id"`
	TypeID        int16       `boil:"type_id" json:"type_id" toml:"type_id" yaml:"type_id"`
	TotalPrice    float64     `boil:"total_price" json:"total_price" toml:"total_price" yaml:"total_price"`
	TotalDiscount float64     `boil:"total_discount" json:"total_discount" toml:"total_discount" yaml:"total_discount"`
	Vat           float64     `boil:"vat" json:"vat" toml:"vat" yaml:"vat"`
	PayableAmount float64     `boil:"payable_amount" json:"payable_amount" toml:"payable_amount" yaml:"payable_amount"`
	TransactionID null.String `boil:"transaction_id" json:"transaction_id,omitempty" toml:"transaction_id" yaml:"transaction_id,omitempty"`
	TimeslotID    string      `boil:"timeslot_id" json:"timeslot_id" toml:"timeslot_id" yaml:"timeslot_id"`
	Data          null.JSON   `boil:"data" json:"data,omitempty" toml:"data" yaml:"data,omitempty"`
	CreatedAt     null.Time   `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt     null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	DeletedAt     null.Time   `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	Discounts     null.JSON   `boil:"discounts" json:"discounts,omitempty" toml:"discounts" yaml:"discounts,omitempty"`
	TimeslotTitle null.String `boil:"timeslot_title" json:"timeslot_title,omitempty" toml:"timeslot_title" yaml:"timeslot_title,omitempty"`
	EventID       null.String `boil:"event_id" json:"event_id,omitempty" toml:"event_id" yaml:"event_id,omitempty"`
	EventTitle    null.String `boil:"event_title" json:"event_title,omitempty" toml:"event_title" yaml:"event_title,omitempty"`
}

type FaktorFullModel struct {
	ID            string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID        string      `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	CartID        string      `boil:"cart_id" json:"cart_id" toml:"cart_id" yaml:"cart_id"`
	Status        int16       `boil:",bind" json:"status_id" toml:"status_id" yaml:"status_id"`
	Type          int16       `boil:",bind" json:"type_id" toml:"type_id" yaml:"type_id"`
	TotalPrice    float64     `boil:"total_price" json:"total_price,omitempty" toml:"total_price" yaml:"total_price"`
	Discounts     null.JSON   `boil:"discounts" json:"discounts,omitempty" toml:"discounts" yaml:"discounts"`
	TotalDiscount float64     `boil:"total_discount" json:"total_discount,omitempty" toml:"total_discount" yaml:"total_discount"`
	Vat           float64     `boil:"vat" json:"vat,omitempty" toml:"vat" yaml:"vat"`
	PayableAmount float64     `boil:"payable_amount" json:"payable_amount,omitempty" toml:"payable_amount" yaml:"payable_amount"`
	TransactionID null.String `boil:"transaction_id" json:"transaction_id,omitempty" toml:"transaction_id" yaml:"transaction_id,omitempty"`
	TimeslotID    string      `boil:"timeslot_id" json:"timeslot_id" toml:"timeslot_id" yaml:"timeslot_id"`
	TimeslotTitle null.String `boil:"timeslot_title" json:"timeslot_title,omitempty" toml:"timeslot_title" yaml:"timeslot_title,omitempty"`
	EventID       null.String `boil:"event_id" json:"event_id,omitempty" toml:"event_id" yaml:"event_id,omitempty"`
	EventTitle    null.String `boil:"event_title" json:"event_title,omitempty" toml:"event_title" yaml:"event_title,omitempty"`
	Data          null.JSON   `boil:"data" json:"data,omitempty" toml:"data" yaml:"data,omitempty"`
	CreatedAt     null.Time   `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
}

//type InvoiceDescription struct {
//	//ID of product or service
//	POSID string `json:"posid"`
//	//description of products or services
//	Description string `json:"description"`
//	//count of products or services
//	Count int `json:"count"`
//	//UOM is brief of Unit of Measure
//	UOM string `json:"uom"`
//	//price od each unit of product or service
//	UnitFee float64 `json:"unit_fee"`
//	//= count * unitfee
//	TotalPrice float64 `json:"total_price"`
//	//gather all discounts applied to this service or product
//	Discounts []Discount `json:"discounts"`
//	//calculated discount
//	Discount float64 `json:"discount"`
//	// = totalprice - discount
//	CalculatedPrice float64 `json:"calculated_price"`
//	//value added tax
//	VAT float64 `json:"vat"`
//	//= calculatedprice + vat
//	TaxAddedPrice float64 `json:"tax_added_price"`
//}

type InvoiceDescription struct {
	SeatID         string  `json:"seat_id"`
	SeatRow        string  `json:"seat_row"`
	SeatNumber     string  `json:"seat_number"`
	ZoneID         string  `json:"zone"`
	ZoneName       string  `json:"zone_name"`
	PriceGroupID   string  `json:"price_group_id"`
	PriceGroupName string  `json:"price_group_name"`
	PricePrice     float64 `json:"price"`
	Discounts      []int64 `json:"discounts"`
	DiscountAmount float64 `json:"discount,omitempty"`
	FinalPrice     float64 `json:"final_price,omitempty"`
}

type InvoiceDescriptionSlice []*InvoiceDescription

type Discount struct {
	ID              string  `json:"id"`
	Description     string  `json:"description"`
	DiscountPercent float64 `json:"discount_percent"`
	DiscountAmount  float64 `json:"discount_amount"`
}

// FaktorForPay Request Faktor send for paying
type FaktorForPay struct {
	ID string `json:"id"`
}

type CostFaktorInvoice struct {
	CampaignID  string  `json:"campaign_id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type LiquidationRequest struct {
	EventID    string  `json:"event_id"`
	EventTitle string  `json:"event_title"`
	Fee        float64 `json:"Fee"`
	UserID     string  `json:"user_id"`
}
