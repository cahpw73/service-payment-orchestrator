package requestModelMiddleware

type RegisterAffiliation struct {
	CriteriaSearchId string            `json:"criteriaSearchId"`
	SearchFields     []SearchField     `json:"searchFields"`
	ServiceCode      string            `json:"serviceCode"`
	Year             string            `json:"year"`
	PersonId         string            `json:"personId"`
	IsTemporal       string            `json:"isTemporal"`
	DataAffiliation  []DataAffiliation `json:"dataAffiliation"`
	AccountNumber    string            `json:"accountNumber"`
	ReferenceName    string            `json:"referenceName"`
}
