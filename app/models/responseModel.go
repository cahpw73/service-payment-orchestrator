package models

type Response[T any] struct {
	ErrCode string `json:"errCode"`
	ErrMsg  string `json:"errMsg"`
	Data    T      `json:"data"`
}
