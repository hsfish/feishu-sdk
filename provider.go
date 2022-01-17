package feishu_sdk

type Provider interface {
	CreateToken() (*AccessToken, error)
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	Expired     int64  `json:"expire"`
}

type TenantProvider interface {
	BatchGetId(emails []string, mobiles []string) (*BatchGetIdResponse, error)
}

type BatchGetIdResponse struct {
	EmailUsers      map[string][]SimpleUser `json:"email_users"`
	EmailsNotExist  []string                `json:"emails_not_exist"`
	MobileUsers     map[string][]SimpleUser `json:"mobile_users"`
	MobilesNotExist []string                `json:"mobiles_not_exist"`
}

type SimpleUser struct {
	OpenId string `json:"open_Id"`
	UserId string `json:"user_id"`
}
