package feishu_sdk


type Storage interface {
	Set(appId, appSecret, accessToken string, expired int64)
	Get(appId, appSecret string ) string
}

var accessTokenStorage Storage

func SetAccessTokenStorage(s Storage) {
	accessTokenStorage = s
}