package accounting

import "time"

type Factor struct {
	//Factor Unique ID
	ID string `json:"id"`
	//issuing date of factor
	IssueDate time.Time `json:"issue_date"`
	//userID
	IssuingFor string `json:"issuing_for"`
	//payment status of factor {enum: pending, paid, expired}
	PaymentStatus string `json:"payment_status"`
	//sum of taxaddedprice of descriptions
	TotalPrices float64 `json:"total_prices"`
	//summary of descriptions discount
	TotalDiscount float64 `json:"total_discount"`
	//total amount of factor for paying (= total price - total discount)
	PayableAmount float64 `json:"payable_amount"`
	//DOPS is json string contain invoiceDescriptionsSlice
	DOPS string `json:"dops"`
	//description of factor
	Description string `json:"description"`
}

type InvoiceDescription struct {
	//ID of product or service
	POSID string `json:"posid"`
	//description of products or services
	Description string `json:"description"`
	//count of products or services
	Count int `json:"count"`
	//UOM is brief of Unit of Measure
	UOM string `json:"uom"`
	//price od each unit of product or service
	UnitFee float64 `json:"unit_fee"`
	//= count * unitfee
	TotalPrice float64 `json:"total_price"`
	//calculated discount
	Discount float64 `json:"discount"`
	// = totalprice - discount
	CalculatedPrice float64 `json:"calculated_price"`
	//value added tax
	VAT float64 `json:"vat"`
	//= calculatedprice + vat
	TaxAddedPrice float64 `json:"tax_added_price"`
}

type InvoiceDescriptionSlice []*InvoiceDescription
