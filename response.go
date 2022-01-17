package feishu_sdk

import "fmt"

type result interface {
	Verify() error
}

type baseResult struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

func (this *baseResult) Verify() error {
	if this.Code != 0 {
		return fmt.Errorf("code:%d  message:%s", this.Code, this.Msg)
	}
	return nil
}

type baseResultWithData struct {
	baseResult
	Data interface{}
}
