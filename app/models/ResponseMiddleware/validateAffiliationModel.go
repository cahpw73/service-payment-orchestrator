package responseTokenMiddleware

type ValidateAffiliationResponseMiddleware struct {
	ServiceCode     int              `json:"serviceCode"`
	DataAffiliation []Affiliation    `json:"dataAffiliation"`
}
