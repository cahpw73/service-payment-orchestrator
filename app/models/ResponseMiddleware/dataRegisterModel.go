package responseTokenMiddleware

type DataRegister struct {
	Label       string  `json:"label"`
	Value       string  `json:"value"`
	Mandatory   *bool   `json:"mandatory"`
	Edit        *bool   `json:"edit"`
	Group       *string `json:"group"`
	Description *string `json:"description"`
	Code        *string `json:"code"`
}
