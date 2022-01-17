package feishu_sdk

type Provider interface {
	CreateToken() (*AccessToken, error)
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	Expired     int64  `json:"expire"`
}
