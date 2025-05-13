package requestModelMiddleware

type ValidateAffiliationRequestMiddleware struct {
	ServiceCode                		int            	`json:"serviceCode"`
	Year                       		int            	`json:"year"`
	SearchCriteriaId           		string         	`json:"searchCriteriaId"`
	SearchCriteriaIdAbbreviation 	string       	`json:"searchCriteriaIdAbbreviation"`
	PersonId                   		int            	`json:"personId"`
	SearchFields               		[]SearchField  	`json:"searchFields"`
}
