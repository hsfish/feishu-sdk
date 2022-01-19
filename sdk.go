package feishu_sdk

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/hsfish/feishu-sdk/util/jsonUtil"
	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
)

func init() {
	http.DefaultClient.Timeout = time.Second * 30
	t := http.DefaultTransport.(*http.Transport)
	t.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

type Sdk struct {
	appId       string
	appSecret   string
	provider    Provider
	contentType string
}

func BuildSdk(appId, appSecret string, opts ...Options) *Sdk {
	c := &Sdk{
		appId:       appId,
		appSecret:   appSecret,
		contentType: "application/json; charset=utf-8",
	}
	for i := range opts {
		opts[i](c)
	}
	return c
}

func (this *Sdk) GetAppId() string {
	return this.appId
}

func (this *Sdk) formatRawQuery(rawQuery map[string]interface{}) string {
	rv := url.Values{}
	for k, v := range rawQuery {
		reflectVal := reflect.ValueOf(v)
		if kind := reflectVal.Kind(); kind != reflect.Array && kind != reflect.Slice {
			rv.Add(k, fmt.Sprintf("%v", v))
			continue
		}
		if n := reflectVal.Len(); n > 0 {
			for i := 0; i < n; i++ {
				rv.Add(k, fmt.Sprintf("%v", reflectVal.Index(i).Interface()))
			}
		} else {
			rv.Add(k, "")
		}
	}
	return rv.Encode()
}

func (this *Sdk) Post(api string, body interface{}, r result) ([]byte, error) {
	return this.send(http.MethodPost, api, nil, body, r)
}

func (this *Sdk) GetWithAuth(api string, rawQuery map[string]interface{}, body interface{}, r result) ([]byte, error) {
	if len(rawQuery) > 0 {
		if strings.Contains(api, "?") {
			api += "&" + this.formatRawQuery(rawQuery)
		} else {
			api += "?" + this.formatRawQuery(rawQuery)
		}
	}
	header, err := this.getTokenHeader()
	if err != nil {
		return nil, err
	}
	return this.send(http.MethodGet, api, header, body, r)
}

func (this *Sdk) PostWithAuth(api string, rawQuery map[string]interface{}, body interface{}, r result) ([]byte, error) {
	if len(rawQuery) > 0 {
		if strings.Contains(api, "?") {
			api += "&" + this.formatRawQuery(rawQuery)
		} else {
			api += "?" + this.formatRawQuery(rawQuery)
		}
	}
	header, err := this.getTokenHeader()
	if err != nil {
		return nil, err
	}
	return this.send(http.MethodPost, api, header, body, r)
}

func (this *Sdk) getTokenHeader() (map[string]string, error) {

	token, err := this.GenToken()
	if err != nil {
		return nil, err
	}
	return map[string]string{"Authorization": "Bearer " + token}, nil
}

func (this *Sdk) formatRequestBody(body interface{}) (io.Reader, error) {
	var reader io.Reader
	if !reflect2.IsNil(body) {
		switch d := body.(type) {
		case string:
			reader = strings.NewReader(d)
		case []byte:
			reader = bytes.NewBuffer(d)
		case io.Reader:
			reader = d
		default:
			if b, err := jsoniter.Marshal(d); err == nil {
				reader = bytes.NewReader(b)
			} else {
				return nil, err
			}
		}
	}
	return reader, nil
}

func (this *Sdk) send(method string, api string, header map[string]string, body interface{}, r result) (buf []byte, err error) {
	reqBody, err := this.formatRequestBody(body)
	if err != nil {
		return nil, err
	}

	if enablePrintln() {
		start := time.Now()
		printInfo("[%d] send from '%s', api: '%s', method: '%s', header: '%s' body: '%s'", start.UnixNano(), this.appId, api, method, jsonUtil.MustMarshalToString(header), jsonUtil.MustMarshalToString(body))
		defer func() {
			printInfo("[%d] recv from '%s', ttl: '%s', err: '%v', body: '%s'", start.UnixNano(), this.appId, time.Since(start), err, string(buf))
		}()
	}

	req, err := http.NewRequest(method, api, reqBody)
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}
	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", this.contentType)
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf(response.Status)
		return
	}

	if buf, err = ioutil.ReadAll(response.Body); err == nil && !reflect2.IsNil(r) {
		if err = jsoniter.Unmarshal(buf, r); err == nil {
			err = r.Verify()
		}
	}
	return
}
