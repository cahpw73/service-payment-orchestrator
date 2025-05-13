package requestModelMiddleware

type SupplementaryData struct {
	AffiliationCode     string `json:"affiliationCode"`
	ServiceCode         string `json:"serviceCode"`
	InvoiceNITCI        string `json:"invoiceNITCI"`
	InvoiceName         string `json:"invoiceName"`
	Company             string `json:"company"`
	IdGeneratedForDebt  string `json:"idGeneratedForDebt"`
	Description         string `json:"description"`
}
