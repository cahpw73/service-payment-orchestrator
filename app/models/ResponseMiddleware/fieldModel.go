package responseTokenMiddleware

type Field struct {
	Identifier      int      `json:"identifier"`
	Label           string   `json:"label"`
	Description     string   `json:"description"`
	Abbreviation    string   `json:"abbreviation"`
	IsMandatory     string   `json:"isMandatory"`
	MinimumLength   string   `json:"minimumLength"`
	MaximumLength   string   `json:"maximumLength"`
	DataTypeCode    string   `json:"dataTypeCode"`
	Values          []Value  `json:"values"`
}
