package mafengwo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/duke-git/lancet/v2/datetime"
	"github.com/duke-git/lancet/v2/netutil"
	"github.com/duke-git/lancet/v2/random"
	"github.com/tiantour/rsae"
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

	body, err := rsae.NewAES().Encrypt(data, key, key[:16])
	if err != nil {
		return nil, err
	}
	base64Data := cryptor.Base64StdEncode(string(body))

	timestamp := datetime.NewUnixNow().ToUnix()
	nonce := random.RandString(16)
	sign := fmt.Sprintf("%d%s%d%s%s%s", PartnerID, action, timestamp, AseKey, nonce, base64Data)
	sign = cryptor.Md5String(sign)

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

	body, err = io.ReadAll(b)
	if err != nil {
		return nil, err
	}

	client := netutil.NewHttpClient()
	resp, err := client.SendRequest(&netutil.HttpRequest{
		RawURL: "https://openapi.mafengwo.cn/deals/rest",
		Method: "POST",
		Body:   body,
		Headers: http.Header{
			"User-Agent":   []string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.67 Safari/537.36"},
			"Content-Type": []string{w.FormDataContentType()},
		},
	})
	if err != nil {
		return nil, err
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	base64Data = cryptor.Base64StdDecode(string(body))
	body, err = rsae.NewAES().Decrypt([]byte(base64Data), key, key[:16])
	if err != nil {
		return nil, err
	}

	result := Deals{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	if result.Error != "" {
		return nil, errors.New(result.Error)
	}

	if result.ErrNo != 1000 {
		return nil, errors.New(result.Message)
	}
	return json.Marshal(result.Data)
}
