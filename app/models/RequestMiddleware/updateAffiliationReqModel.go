package requestModelMiddleware

type UpdateAffiliationRequestMiddleware struct {
	PersonId         string `json:"personId"`
	AffiliationCode  string `json:"affiliationCode"`
	StateAffiliation string `json:"stateAffiliation"`
}
