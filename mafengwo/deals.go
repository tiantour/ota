package mafengwo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"

	"github.com/tiantour/fetch"
	"github.com/tiantour/imago"
	"github.com/tiantour/rsae"
	"github.com/tiantour/tempo"
)

// Deals deals
type Deals struct {
	ErrNo   int32                  `json:"errno"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
	Error   string                 `json:"error"`
}

// NewDeals new deals
func NewDeals() *Deals {
	return &Deals{}
}

// Fetch fetch rest
func (d *Deals) Fetch(action string, data []byte) ([]byte, error) {
	key := []byte(AseKey)
	iv := []byte(AseKey)[:16]

	aesData, err := rsae.NewAES().Encrypt(data, key, iv)
	if err != nil {
		return nil, err
	}
	base64Data := rsae.NewBase64().Encode(aesData)

	timestamp := tempo.NewNow().Unix()
	nonce := imago.NewRandom().Text(16)
	sign := fmt.Sprintf("%d%s%d%s%s%s", PartnerID, action, timestamp, AseKey, nonce, base64Data)
	sign = rsae.NewMD5().Encode(sign)

	b := &bytes.Buffer{}
	w := multipart.NewWriter(b)
	w.WriteField("partnerId", fmt.Sprintf("%d", PartnerID))
	w.WriteField("action", action)
	w.WriteField("timestamp", fmt.Sprintf("%d", timestamp))
	w.WriteField("nonce", nonce)
	w.WriteField("data", base64Data)
	w.WriteField("sign", sign)
	w.WriteField("access_token", AccessToken)
	w.WriteField("file_data", "")
	w.Close()

	body, err := ioutil.ReadAll(b)
	if err != nil {
		return nil, err
	}
	body, err = fetch.Cmd(fetch.Request{
		Method: "POST",
		URL:    "https://openapi.mafengwo.cn/deals/rest",
		Body:   body,
		Header: http.Header{
			"User-Agent":   []string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.67 Safari/537.36"},
			"Content-Type": []string{w.FormDataContentType()},
		},
	})
	if err != nil {
		return nil, err
	}

	result := Deals{}
	if len(body) < 100 {
		err = json.Unmarshal(body, &result)
		if err != nil {
			return nil, err
		}
		if result.Error != "" {
			return nil, errors.New(result.Error)
		}
	}

	body, err = rsae.NewBase64().Decode(string(body))
	if err != nil {
		return nil, err
	}
	body, err = rsae.NewAES().Decrypt(body, key, iv)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	if result.ErrNo != 1000 {
		return nil, errors.New(result.Message)
	}
	return json.Marshal(result.Data)
}
