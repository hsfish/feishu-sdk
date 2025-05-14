package feishu_sdk

import (
	"testing"

	"github.com/hsfish/feishu-sdk/util/jsonUtil"
	"github.com/stretchr/testify/assert"
)

func TestSdk_SendMessageMulti(t *testing.T) {
	result, err := testSdk.SendMessageMulti(&UserIdArgs{
		UserIds: []string{"a2fd43gd"},
	}, &TextMessage{Content: "推送测试"})
	if assert.NoError(t, err) {
		t.Log(jsonUtil.MustMarshalToString(result))
	}
}

func TestSdk_SendMessageMulti_Post(t *testing.T) {

	zhCn := &PostMessageContent{
		Title: "hsfish_test",
		Content: PostMessageTagList{}.Append(
			NewPostMessageTag().AddTextTag("第一行：").AddTextTag("测试Text\n测试Text2"),
			NewPostMessageTag().AddTextTag("第二行：").AddATag("这是一个超链接", "http://www.baidu.com"),
		),
	}

	result, err := testSdk.SendMessageMulti(&UserIdArgs{
		UserIds: []string{"a2fd43gd"},
	}, &PostMessage{
		ZhCn: zhCn,
	})
	if assert.NoError(t, err) {
		t.Log(jsonUtil.MustMarshalToString(result))
	}
}

type a struct {
}

func (this a) GetType() string {
	return ""
}

func TestSdk_SendMessageMulti_Invalid(t *testing.T) {
	_, err := testSdk.SendMessageMulti(&UserIdArgs{
		UserIds: []string{"a2fd43gd"},
	}, a{})
	assert.Error(t, err)
}

func TestSdk_SendMessage_Text(t *testing.T) {
	result, err := testSdk.SendMessage(UserIdType_Chat_Id, "oc_0413ff780d2c02e945cc671cb9562bd2", &TextMessage{Content: "推送测试"})
	if assert.NoError(t, err) {
		t.Log(result.MessageId)
	}
}

func TestSdk_SendMessage_Post(t *testing.T) {
	result, err := testSdk.SendMessage(UserIdType_Chat_Id, "oc_0413ff780d2c02e945cc671cb9562bd2",
		&PostMessage{
			ZhCn: &PostMessageContent{
				Title: "hsfish_test",
				Content: PostMessageTagList{}.Append(
					NewPostMessageTag().AddTextTag("第一行：").AddTextTag("测试Text\n测试Text2"),
					NewPostMessageTag().AddTextTag("第二行：").AddATag("这是一个超链接", "http://www.baidu.com"),
				),
			},
		},
	)
	if assert.NoError(t, err) {
		t.Log(result.MessageId)
	}
}

func TestSdk_SendMessage_Interactive(t *testing.T) {

	result, err := testSdk.SendMessage(UserIdType_Chat_Id, "oc_0413ff780d2c02e945cc671cb9562bd2",
		NewSimpleInteractiveMessage("test", "testHahaHa"),
	)
	if assert.NoError(t, err) {
		t.Log(result.MessageId)
	}

	result, err = testSdk.SendMessage(UserIdType_Chat_Id, "oc_0413ff780d2c02e945cc671cb9562bd2",
		NewLinkInteractiveMessage("test", "testHahaHa", "https://www.baidu.com"),
	)
	if assert.NoError(t, err) {
		t.Log(result.MessageId)
	}
}
