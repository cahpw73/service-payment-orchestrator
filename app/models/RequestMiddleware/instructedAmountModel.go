package requestModelMiddleware

type InstructedAmount struct {
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
}
