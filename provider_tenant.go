package feishu_sdk

const (
	api_Auth_TenantAccessToken_Internal_V3 = "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal"
)

type TenantProvider struct {
	*Sdk
}

func (this *TenantProvider) CreateToken() (*AccessToken, error) {
	resp := &struct {
		baseResult
		AccessToken string `json:"tenant_access_token"`
		Expired     int64  `json:"expire"`
	}{}
	if _, err := this.Post(api_Auth_TenantAccessToken_Internal_V3, map[string]interface{}{
		"app_id":     this.appId,
		"app_secret": this.appSecret,
	}, resp); err != nil {
		return nil, err
	}

	return &AccessToken{AccessToken: resp.AccessToken, Expired: resp.Expired}, nil
}
