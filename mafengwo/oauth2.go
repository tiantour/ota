package mafengwo

import (
	"encoding/json"
	"fmt"

	"github.com/tiantour/fetch"
)

// Oauth2 oauth2
type Oauth2 struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

// NewOauth2 new oauth2
func NewOauth2() *Oauth2 {
	return &Oauth2{}
}

// Token get access token
func (o *Oauth2) Token() (*Oauth2, error) {
	path := fmt.Sprintf("https://openapi.mafengwo.cn/oauth2/token?grant_type=client_credentials&client_id=%d&client_secret=%s", ClientID, ClientSecret)
	body, err := fetch.Cmd(&fetch.Request{
		Method: "GET",
		URL:    path,
	})
	if err != nil {
		return nil, err
	}
	return o, json.Unmarshal(body, o)
}
