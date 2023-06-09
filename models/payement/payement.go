package payement

type Checkouts struct {
	Gateways     []Gateway `json:"gateways"`
	WalletCharge float64   `json:"wallet_charge"`
}

type Gateway struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
	ID   string `json:"id"`
}
