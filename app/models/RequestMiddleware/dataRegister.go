package requestModelMiddleware

type DataRegister struct {
	Code        string `json:"code"`
	Edit        string `json:"edit"`
	Description string `json:"description"`
	Label       string `json:"label"`
	Value       string `json:"value"`
	Mandatory   string `json:"mandatory"`
	Group       string `json:"group"`
}
