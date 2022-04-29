package model

type ResponseLogin struct {
	TokenType   string `json:"tokenType"`
	AccessToken string `json:"accessToken"`
	ExpiredAt   int64  `json:"expiredAt"`
}
