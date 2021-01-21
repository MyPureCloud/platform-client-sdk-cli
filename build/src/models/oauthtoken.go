package models

import (
	"encoding/json"
)

type OAuthToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Error       string `json:"error"`
}

type OAuthTokenData struct {
	OAuthToken
	OAuthTokenExpiry string `json:"oauth_token_expiry"`
}

func (d OAuthTokenData) String() string {
	bytes, _ := json.Marshal(d)
	return string(bytes)
}