package feishu_sdk

import (
	"fmt"
	"reflect"

	"github.com/hsfish/feishu-sdk/util/jsonUtil"
)

const (
	api_Message_BatchSend_V4 = "https://open.feishu.cn/open-apis/message/v4/batch_send/"
	api_Message_Send_V1      = "https://open.feishu.cn/open-apis/im/v1/messages"
)

const (
	MessageType_Text        = "text"        // 文本
	MessageType_Image       = "image"       // 图片
	MessageType_Post        = "post"        // 富文本
	MessageType_ShareChat   = "share_chat"  // 群名片
	MessageType_Interactive = "interactive" // 卡片
)

type Message interface {
	GetType() string
}

type MessageTypeContent interface {
	GetContent() interface{}
}

type MessageTypeCard interface {
	GetCard() interface{}
}

// 文本消息
type TextMessage struct {
	Content string
}

func (this *TextMessage) GetType() string {
	return MessageType_Text
}

func (this *TextMessage) GetContent() interface{} {
	return map[string]interface{}{
		"text": this.Content,
	}
}

// 图片消息
type ImageMessage struct {
	Content string
}

func (this *ImageMessage) GetType() string {
	return MessageType_Image
}

func (this *ImageMessage) GetContent() interface{} {
	return map[string]interface{}{
		"image_key": this.Content,
	}
}

// 群名片
type ShareChatMessage struct {
	Content string
}

func (this *ShareChatMessage) GetType() string {
	return MessageType_ShareChat
}

func (this *ShareChatMessage) GetContent() interface{} {
	return map[string]interface{}{
		"share_chat_id": this.Content,
	}
}

// 富文本消息
type PostMessage struct {
	ZhCn *PostMessageContent `json:"zh_cn,omitempty"`
	EnUs *PostMessageContent `json:"en_us,omitempty"`
	JaJp *PostMessageContent `json:"ja_jp,omitempty"`
}

func (this *PostMessage) GetType() string {
	return MessageType_Post
}

func (this *PostMessage) GetContent() interface{} {
	return map[string]interface{}{
		"post": this,
	}
}

// 富文本内容
type PostMessageContent struct {
	Title   string             `json:"title"`
	Content PostMessageTagList `json:"content"`
}

type PostMessageTagList []PostMessageTag

func (this PostMessageTagList) Append(tags ...PostMessageTag) PostMessageTagList {
	this = append(this, tags...)
	return this
}

type PostMessageTag []map[string]interface{}

func NewPostMessageTag() PostMessageTag {
	return PostMessageTag{}
}

// unEscape. unescape解码
func (this PostMessageTag) AddTextTag(text string, unEscape ...bool) PostMessageTag {
	this = append(this, map[string]interface{}{
		"tag":       "text",
		"text":      text,
		"un_escape": append(unEscape, false)[0],
	})
	return this
}

// text.文本内容
// href.超链接
func (this PostMessageTag) AddATag(text string, href string) PostMessageTag {
	return append(this, map[string]interface{}{
		"tag":  "a",
		"text": text,
		"href": href,
	})
}

// userName.用户姓名
func (this PostMessageTag) AddAtTag(openId, userName string) PostMessageTag {
	return append(this, map[string]interface{}{
		"tag":       "at",
		"user_id":   openId,
		"user_name": userName,
	})
}

// imageKey.图片的唯一标识
// height.高度
// width.宽度
func (this PostMessageTag) AddImgTag(imageKey string, height, width int) PostMessageTag {
	return append(this, map[string]interface{}{
		"tag":       "img",
		"image_key": imageKey,
		"height":    height,
		"width":     width,
	})
}

// 卡片消息
type InteractiveMessage struct {
	Config       *CardConfig     `json:"config,omitempty"`
	CardLink     *CardElementUrl `json:"card_link,omitempty"`
	Header       *CardHeader     `json:"header,omitempty"`
	Elements     []interface{}   `json:"elements"`
	I18nElements *I18nElement    `json:"i18n_elements,omitempty"`
}

func (this *InteractiveMessage) GetType() string {
	return MessageType_Interactive
}

func (this *InteractiveMessage) GetCard() interface{} {
	return this
}

type CardConfig struct {
	WideScreenMode bool `json:"wide_screen_mode"`
	EnableForward  bool `json:"enable_forward"`
}

type CardElementUrl struct {
	Url        string `json:"url"`
	AndroidUrl string `json:"android_url"`
	IosUrl     string `json:"ios_url"`
	PcUrl      string `json:"pc_url"`
}

type CardHeader struct {
	Title    *CardHeaderTitle `json:"title,omitempty"`
	Template string           `json:"template"`
}

type CardHeaderTitle struct {
	Tag     string    `json:"tag"`
	Content string    `json:"content"`
	Lines   int       `json:"lines,omitempty"`
	I18n    *CardI18n `json:"i18n,omitempty"`
}

type CardI18n struct {
	ZhCn string `json:"zh_cn"`
	EnUs string `json:"en_us"`
	JaJp string `json:"ja_jp"`
}

type I18nElement struct {
	ZhCn []interface{} `json:"zh_cn"`
	EnUs []interface{} `json:"en_us"`
	JaJp []interface{} `json:"ja_jp"`
}

type SendMessageMultiResult struct {
	MessageId            string   `json:"message_id"`
	InvalidDepartmentIds []string `json:"invalid_department_ids"`
	InvalidOpenIds       []string `json:"invalid_open_ids"`
	InvalidUserIds       []string `json:"invalid_user_ids"`
}

func (this *Sdk) SendMessageMulti(users *UserIdArgs, msg Message) (*SendMessageMultiResult, error) {
	result := &baseResultWithData{Data: &SendMessageMultiResult{}}
	body := users.GetData()
	body["msg_type"] = msg.GetType()
	switch obj := msg.(type) {
	case MessageTypeContent:
		body["content"] = obj.GetContent()
	case MessageTypeCard:
		body["card"] = obj.GetCard()
	default:
		return nil, fmt.Errorf("暂不支持该类型：%s %s", reflect.TypeOf(msg).Name(), reflect.TypeOf(msg).Kind())
	}

	_, err := this.PostWithAuth(api_Message_BatchSend_V4, nil, body, result)
	if err != nil {
		return nil, err
	}
	return result.Data.(*SendMessageMultiResult), nil
}

type SendMessageResult struct {
	MessageId string `json:"message_id"`
}

func (this *Sdk) SendMessage(recvIdType UserIdType, recvId string, msg Message) (*SendMessageResult, error) {
	result := &baseResultWithData{Data: &SendMessageResult{}}
	body := map[string]interface{}{
		"receive_id": recvId,
		"msg_type":   msg.GetType(),
	}
	switch obj := msg.(type) {
	case *PostMessage:
		body["content"] = jsonUtil.MustMarshalToString(obj)
	case MessageTypeContent:
		body["content"] = jsonUtil.MustMarshalToString(obj.GetContent())
	case MessageTypeCard:
		body["content"] = jsonUtil.MustMarshalToString(obj.GetCard())
	default:
		return nil, fmt.Errorf("暂不支持该类型：%s %s", reflect.TypeOf(msg).Name(), reflect.TypeOf(msg).Kind())
	}

	if _, err := this.PostWithAuth(api_Message_Send_V1, map[string]interface{}{
		"receive_id_type": recvIdType,
	}, body, result); err != nil {
		return nil, err
	}
	return result.Data.(*SendMessageResult), nil
}
