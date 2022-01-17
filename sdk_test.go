package feishu_sdk

import (
	"github.com/stretchr/testify/assert"
	"hsfish/feishu-sdk/util/jsonUtil"
	"testing"
)

func TestMain(m *testing.M) {
	SetDefaultLogger(&consoleLogger{})
	m.Run()
}

var testSdk = BuildSdk("cli_a15c61d67038900b", "gdP9ucfwybXA9IGAXdiEIc60SzIDjPz3", WithTenantProvider())

func TestSdk_BatchGetUserId(t *testing.T) {
	// a2fd43gd
	resp, err := testSdk.BatchGetUserId(nil, []string{"18575581607"})
	if assert.NoError(t, err) {
		t.Log(jsonUtil.MustMarshalToString(resp))
	}
}
