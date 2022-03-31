package models

type Token struct {
	Token      string `json:"accessToken"`
	Username   string `json:"username"`
	ExpireTime int64  `json:"expire"`
}
