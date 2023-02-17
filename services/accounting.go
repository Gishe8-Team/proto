package services

const (
	CreateBuyFactor     = "9.1.1"
	SendFactorToPayment = "9.1.2"
	CreateCostFactor    = "9.1.3"
)

const (
	FindFactorByProfileID  = "9.2.1"
	FindFactorsByEventID   = "9.2.2"
	FindFactorsByTimestamp = "9.2.3"
	FindUnpaidFactors      = "9.2.4"
)

const (
	CalculateEventProfit         = "9.3.1"
	ListEventFactors             = "9.3.2"
	CalculateEventSharing        = "9.3.3"
	CreateLiquidationTransaction = "9.3.4"
)

const (
	CreateCoupon  = "9.4.1"
	EditCoupon    = "9.4.2"
	DeleteCoupon  = "9.4.3"
	ListCoupons   = "9.4.4"
	FindCoupons   = "9.4.5"
	CreateSubsidy = "9.4.6"
	EditSubsidy   = "9.4.7"
	DeleteSubsidy = "9.4.8"
	ListSubsidies = "9.4.9"
	FindSubsidies = "9.4.10"
)
