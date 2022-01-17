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
		expired: time.Now().Add(time.Duration(expired-100) * time.Second),
	})
}

func (this *memoryStorage) Get(appId string) string {
	unit, ok := this.data.Load(appId)
	if ok && unit.(*storageUnit).expired.Before(time.Now()) {
		return unit.(*storageUnit).token
	}
	return ""
}

type storageUnit struct {
	token   string
	expired time.Time
}
