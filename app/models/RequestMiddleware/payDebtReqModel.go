package requestModelMiddleware

type PayDebtRequestMiddleware struct {
	OwnerAccount     	OwnerAccount     	`json:"ownerAccount"`
	DebtorAccount    	DebtorAccount    	`json:"debtorAccount"`
	InstructedAmount 	InstructedAmount 	`json:"instructedAmount"`
	SupplementaryData 	SupplementaryData 	`json:"supplementaryData"`
	Risk             	Risk             	`json:"risk"`
}
