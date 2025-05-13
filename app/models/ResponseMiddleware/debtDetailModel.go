package responseTokenMiddleware

type DebtDetail struct {
	Description           	string  `json:"description"`
	ReferenceCode         	string  `json:"referenceCode"`
	MonthPeriod           	int     `json:"monthPeriod"`
	YearPeriod            	int     `json:"yearPeriod"`
	CommissionAmount      	float64 `json:"commissionAmount"`
	CurrencyCode          	string  `json:"currencyCode"`
	Amount                	float64 `json:"amount"`
	AccumulatedAmount     	float64 `json:"accumulatedAmount"`
	Identifier            	int     `json:"identifier"`
	ValidationType        	string  `json:"validationType"`
	Detail                	string  `json:"detail"`
	AdditionalDataDetails 	string  `json:"additionalDataDetails"`
	PaymentPlanCode       	*string `json:"paymentPlanCode"`
	IdGeneratedForDebt     	string  `json:"idGeneratedForDebt"`
}
