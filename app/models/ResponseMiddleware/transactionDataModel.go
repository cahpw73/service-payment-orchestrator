package responseTokenMiddleware

type TransactionDataResponseMiddleware struct {
	Status         string         `json:"status"`
	MAEId          string         `json:"maeId"`
	NroTransaction string         `json:"nroTransaction"`
	ReceiptDetail  ReceiptDetail  `json:"receiptDetail"`
}
