package services

const (
	CreateBookingSlot        = "8.1.1" // incoming message from event when starting time =~ -5 minutes since now to create a Meta and seats for it if booking mode is Layout
	ChangeBookingSlotStatus  = "8.1.2" // incoming message from event when only status os changed during booking phase
	ChangeBookingPriceGroups = "8.1.3" // incoming message from event when only price groups changed during booking this need to send a message to WS service for changing Prices on client side real-time
	ChangeBookingSeats       = "8.1.4" // incoming message from event when booking only seats changed while booking is in progress this one also needs to update client side booking with WS message
	FinalizeBookingSlot      = "8.1.5" // incoming message from event when booking in end time is = time now() and we need to complete all the carts etc...
)

const (
	ViewCart         = "8.2.1" // RequestViewCartModel -> ResponseViewCartModelLayout & NoLayout
	SubmitCart       = "8.2.2" // RequestSubmitCartModel -> Submit Cart generate a message for marketing and accounting
	ChangeCartStatus = "8.2.3" // ChangeCartStatus REQ Body incoming request from payment to notify booking that payment status is done this needs to send a message to accounting to notify it to
)

const (
	GetSingleBookingSlot = "8.3.1" // this is for admin usage

)

const (
	ToggleCartSeat = "8.66.1"
)
