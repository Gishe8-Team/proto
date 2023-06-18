package services

const (
	CreateBuyFactor         = "9.1.1"
	SendFactorToPayment     = "9.1.2"
	CreateCostFactor        = "9.1.3"
	CreateLiquidationFactor = "9.1.4"
)

const (
	FindFactorWithID       = "9.2.1"
	FindFactorByProfileID  = "9.2.2"
	FindFactorsByEventID   = "9.2.3"
	FindFactorsByTimestamp = "9.2.4"
	FindUnpaidFactors      = "9.2.5"
)

const (
	CalculateEventProfit         = "9.3.1"
	CalculateEventSharing        = "9.3.2"
	CreateLiquidationTransaction = "9.3.3"
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
