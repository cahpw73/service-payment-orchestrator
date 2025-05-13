package responseTokenMiddleware

type SearchCriteriaResponseMiddleware struct {
	ServiceCode     string          	`json:"serviceCode"`
	Year            *int            	`json:"year"`
	SubServices     []SubService    	`json:"subServices"`
	SearchCriteria  []SearchCriteria 	`json:"searchCriteria"`
}
