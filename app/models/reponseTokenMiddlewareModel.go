package models

type ResponseTokenMiddleware struct {
	AccessToken string `json:"accessToken"`
	TokenType   string `json:"tokenType"`
	ExpiresIn   int    `json:"expiresIn"`
	Scope       string `json:"scope"`
	Jti         string `json:"jti"`
}
