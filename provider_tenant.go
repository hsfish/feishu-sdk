package feishu_sdk

func BuildTenantProvider(appId, appSecret string) Provider {
	return &Tenant{sdk: buildSdk(appId, appSecret)}
}

const (
	api_Tenant_AccessToken_Internal = "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal"
	api_Tenant_User_V1_Batch_Get_Id = "https://open.feishu.cn/open-apis/user/v1/batch_get_id"
)

type Tenant struct {
	*sdk
}

func (this *Tenant) CreateToken() (*AccessToken, error) {
	resp := &struct {
		baseResponse
		AccessToken string `json:"tenant_access_token"`
		Expired     int64  `json:"expire"`
	}{}
	if _, err := this.Post(api_Tenant_AccessToken_Internal, map[string]interface{}{
		"app_id":     this.appId,
		"app_secret": this.appSecret,
	}, resp); err != nil {
		return nil, err
	}

	return &AccessToken{AccessToken: resp.AccessToken, Expired: resp.Expired}, nil
}

func (this *Tenant) BatchGetId(emails []string, mobiles []string) (*BatchGetIdResponse, error) {
	resp := &struct {
		baseResponse
		*BatchGetIdResponse
	}{}
	if _, err := this.Post(api_Tenant_AccessToken_Internal, map[string]interface{}{
		"app_id":     this.appId,
		"app_secret": this.appSecret,
	}, resp); err != nil {
		return nil, err
	}

	return &AccessToken{AccessToken: resp.AccessToken, Expired: resp.Expired}, nil
}
