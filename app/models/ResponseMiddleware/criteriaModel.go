package responseTokenMiddleware

type SearchCriteria struct {
	LabelCriteria            		string  `json:"labelCriteria"`
	Description              		string  `json:"description"`
	SearchCriteriaId         		string  `json:"searchCriteriaId"`
	SearchCriteriaIdAbbreviation 	string  `json:"searchCriteriaIdAbbreviation"`
	Fields                   		[]Field `json:"fields"`
}
