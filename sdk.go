package feishu_sdk

import (
	"bytes"
	"crypto/tls"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func init() {
	http.DefaultClient.Timeout = time.Second * 30
	t := http.DefaultClient.Transport.(*http.Transport)
	t.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

type sdk struct {
	appId     string
	appSecret string
}

type Response interface {
	Verify() error
}

type baseResponse struct {
	Code int64  `json:"code"`
	Msg  string `json:"message"`
}

func (this *baseResponse) Verify() error {
	if this.Code != 0 {
		return fmt.Errorf("code:%d  message:%s", this.Code, this.Msg)
	}
	return nil
}

func (this *sdk) send(method string, api string, header map[string]string, body interface{}, resp Response) ([]byte, error) {
	var r io.Reader
	if reflect2.IsNil(body) {
		switch d := body.(type) {
		case string:
			r = strings.NewReader(d)
		case []byte:
			r = bytes.NewBuffer(d)
		case io.Reader:
			r = d
		default:
			if b, err := jsoniter.Marshal(d); err != nil {
				return nil, err
			} else {
				r = bytes.NewReader(b)
			}
		}
	}

	req, err := http.NewRequest(method, api, r)
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusSeeOther {
		return nil, fmt.Errorf(response.Status)
	}
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if reflect2.IsNil(resp) {
		if err = jsoniter.Unmarshal(b, resp); err != nil {
			return nil, err
		}
		return b, resp.Verify()
	}

	return b, nil

}
