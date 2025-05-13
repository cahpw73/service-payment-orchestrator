package requestModelMiddleware

type GetDebtsRequestMiddleware struct {
	ServiceCode     int    `json:"serviceCode"`
    PersonID    	int    `json:"personId"`
    Year            int    `json:"year"`
    AffiliationCode string `json:"affiliationCode"`
}
