package booking

type CartStatus int8

const (
	CartOpen CartStatus = iota + 1
	CartPending
	CartFinalized
)
