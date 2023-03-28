package types

type PriceRes struct {
	Ticker string  `json:"ticker"`
	Price  float64 `json:"price"`
}
