package responseTokenMiddleware

type GetDebtsResponseMiddleware struct {
	AffiliationCode       string       `json:"affiliationCode"`
	ServiceCode           int          `json:"serviceCode"`
	InvoiceTaxId          int          `json:"invoiceTaxId"`
	InvoiceName           *string      `json:"invoiceName"`
	InvoiceCanModifyData  *bool        `json:"invoiceCanModifyData"`
	DebtDetails           []DebtDetail `json:"debtDetails"`
}
