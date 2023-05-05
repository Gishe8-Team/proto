package booking

type TimeslotStatus int

const (
	TimeslotOpen TimeslotStatus = iota + 1
	TimeslotPaused
	TimeslotFinalized
)

type CartStatus int8

const (
	CartOpen CartStatus = iota + 1
	CartPending
	CartFinalized
)
