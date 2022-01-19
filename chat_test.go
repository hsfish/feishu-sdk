package feishu_sdk

import (
	"testing"

	"github.com/hsfish/feishu-sdk/util/jsonUtil"
)

func TestSdk_QueryChat(t *testing.T) {
	result, err := testSdk.QueryChat("", "", 0, true)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result.HasMore, result.PageToken)
	for _, item := range result.List {
		t.Log(jsonUtil.MustMarshalToString(item))
	}
}

func TestSdk_SearchChat(t *testing.T) {
	result, err := testSdk.SearchChat("", "智象", "", 0, true)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result.HasMore, result.PageToken)
	for _, item := range result.List {
		t.Log(jsonUtil.MustMarshalToString(item))
	}
}
