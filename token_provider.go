package feishu_sdk

type Provider interface {
	CreateToken() (*AccessToken, error)
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	Expired     int64  `json:"expire"`
}

func (this *Sdk) GenToken() (string, error) {

	if token := tokenStorage.Get(this.appId); token != "" {
		return token, nil
	}

	if t, err := this.CreateToken(); err != nil {
		return "", err
	} else {
		return t.AccessToken, nil
	}
}

func (this *Sdk) CreateToken() (*AccessToken, error) {
	if t, err := this.provider.CreateToken(); err != nil {
		return nil, err
	} else {
		tokenStorage.Set(this.appId, t.AccessToken, t.Expired)
		return t, nil
	}
}
