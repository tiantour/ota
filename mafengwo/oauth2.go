package mafengwo

import (
	"fmt"

	"github.com/duke-git/lancet/v2/netutil"
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
	client := netutil.NewHttpClient()
	resp, err := client.SendRequest(&netutil.HttpRequest{
		RawURL: fmt.Sprintf("https://openapi.mafengwo.cn/oauth2/token?grant_type=client_credentials&client_id=%d&client_secret=%s", ClientID, ClientSecret),
		Method: "GET",
	})
	if err != nil {
		return nil, err
	}

	result := Oauth2{}
	err = client.DecodeResponse(resp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
