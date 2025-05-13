package requestModelMiddleware

type UpdateAffiliation struct {
	AffiliationCode  string `json:"affiliationCode"`
	AccountNumber    string `json:"accountNumber"`
	StateAffiliation string `json:"stateAffiliation"`
}
