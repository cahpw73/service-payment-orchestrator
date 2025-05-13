package responseTokenMiddleware

type SubCategoryResponseMiddleware struct {
	SubCategoryCod string `json:"subCategoryCod"`
	CategoryCod    string `json:"categoryCod"`
	Description    string `json:"description"`
	HasCity        string `json:"hasCity"`
}
