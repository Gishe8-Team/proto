package booking

type TimeSlotStatus int

const (
	TimeslotStart TimeSlotStatus = iota + 1
	TimeslotPause
	TimeslotStop
	TimeslotFinalize
)

type CartStatus int8

const (
	CartOpen CartStatus = iota + 1
	CartPending
	CartFinalized
)
