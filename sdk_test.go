package feishu_sdk

import (
	"testing"
	"time"

	"gitee.com/hsfish/feishu-sdk/util/jsonUtil"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	SetDefaultLogger(&consoleLogger{})
	m.Run()
}

var testSdk = BuildSdk("", "", WithTenantProvider())

func TestSdk_BatchGetUserId(t *testing.T) {
	// a2fd43gd
	resp, err := testSdk.UserBatchGetId(nil, []string{"18575581607"})
	if assert.NoError(t, err) {
		t.Log(jsonUtil.MustMarshalToString(resp))
	}
}

func TestSdk_GenToken(t *testing.T) {
	tokenStorage.Set(testSdk.GetAppId(), "hsfish_test", 5)
	token, err := testSdk.GenToken()
	if assert.NoError(t, err) && assert.Equal(t, token, "hsfish_test") {
		t.Log("token storage is work：", token)
	}

	time.Sleep(time.Second * 6)
	token, err = testSdk.GenToken()
	if assert.NoError(t, err) && assert.NotEqual(t, token, "hsfish_test") {
		t.Log("token storage flush is work：", token)
	}
}
