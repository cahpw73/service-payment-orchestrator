package requestModelMiddleware

type DataAffiliation struct {
	NameOwner    string         `json:"nameOwner"`
	Code         string         `json:"code"`
	Identify     int            `json:"identify"`
	DataRegister []DataRegister `json:"dataRegister"`
	Description  string         `json:"description"`
}
