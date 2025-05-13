package responseTokenMiddleware

type AffiliationDataResponseMiddleware struct {
	ServiceCode        string `json:"serviceCode"`
	ServiceDesc        string `json:"serviceDesc"`
	ReferenceName      string `json:"referenceName"`
	NameHolder         string `json:"nameHolder"`
	AffiliationCode    string `json:"affiliationCode"`
	InternalCode       string `json:"internalCod"`
	Year               string `json:"year"`
	DescriptionTag     string `json:"descriptionTag"`
	StateContingency   string `json:"stateContingency"`
	MessageContingency string `json:"msjContingency"`
}
