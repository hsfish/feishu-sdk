package feishu_sdk

import (
	"sync"
	"time"
)

type Storage interface {
	Set(appId, accessToken string, expired int64)
	Get(appId string) string
}

var tokenStorage Storage = &memoryStorage{data: sync.Map{}}

func SetDefaultTokenStorage(s Storage) {
	tokenStorage = s
}

type memoryStorage struct {
	data sync.Map
}

func (this *memoryStorage) Set(appId, accessToken string, expired int64) {
	this.data.Store(appId, &storageUnit{
		token:   accessToken,
		expired: time.Now().Add(time.Duration(expired) * time.Second).Unix(),
	})
}

func (this *memoryStorage) Get(appId string) string {
	if unit, ok := this.data.Load(appId); ok {
		if su := unit.(*storageUnit); su.expired-time.Now().Unix() > 10 {
			return su.token
		}
	}
	return ""
}

type storageUnit struct {
	token   string
	expired int64
}
