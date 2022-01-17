package feishu_sdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	SetDefaultLogger(&consoleLogger{})
	m.Run()
}

var testClient = BuildTenantProvider("cli_a15c61d67038900b", "gdP9ucfwybXA9IGAXdiEIc60SzIDjPz3")

func TestTenant_CreateToken(t *testing.T) {
	token, err := testClient.CreateToken()
	if assert.NoError(t, err) {
		t.Log(token.AccessToken, token.Expired)
	}
}
