package responseTokenMiddleware

type ReceiptDetail struct {
	AccountingDate      string  `json:"accountingDate"`
	AccountingTime      string  `json:"accountingTime"`
	AccountingEntry     string  `json:"accountingEntry"`
	FromAccountNumber   string  `json:"fromAccountNumber"`
	FromHolder          string  `json:"fromHolder"`
	Amount              float64 `json:"amount"`
	Currency            string  `json:"currency"`
	ExchangeAmount      float64 `json:"exchangeAmount"`
	FromAccountCurrency string  `json:"fromAccountCurrency"`
	ExchangeRateDebit   string  `json:"exchangeRateDebit"`
	Company             string  `json:"company"`
	AffiliationNumber   string  `json:"affiliationNumber"`
	Description         string  `json:"description"`
	ServicePaymentCode  string  `json:"servicePaymentCode"`
	VoucherId           *string `json:"voucherId"`
}
