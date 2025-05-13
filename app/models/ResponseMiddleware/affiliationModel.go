package responseTokenMiddleware

type Affiliation struct {
	Identify       int            `json:"identify"`
	NameOwner      string         `json:"nameOwner"`
	Code           string         `json:"code"`
	Description    string         `json:"description"`
	AdditionalFact *string        `json:"additionalFact"`
	DataRegister   []DataRegister `json:"dataRegister"`
}
