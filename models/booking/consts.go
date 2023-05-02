package booking

type TimeSlotStatus int

const (
	Open TimeSlotStatus = iota + 1
	Paused
	Finalized
)

type CartStatus int8

const (
	CartOpen CartStatus = iota + 1
	CartPending
	CartFinalized
)
