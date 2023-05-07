package services

const (
	CreateBookingSlot        = "8.1.1"
	ChangeBookingSlotStatus  = "8.1.2"
	ChangeBookingPriceGroups = "8.1.3"
	ChangeBookingSeats       = "8.1.4"
	FinalizeBookingSlot      = "8.1.5" // from 811 to this generate from event
)

const (
	ViewCart       = "8.2.1" // RequestViewCartModel -> ResponseViewCartModelLayout & NoLayout
	FinalizingCart = "8.2.2" // RequestFinalizeCartModel -> Finalizing Cart generate a message for marketing and accounting
)

const (
	GetSingleBookingSlot = "8.3.1" // this is for admin usage

)

const (
	ToggleCartSeat = "8.66.1"
)
