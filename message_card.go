package feishu_sdk

// 卡片消息
type InteractiveMessage struct {
	Config       *CardConfig     `json:"config,omitempty"`
	CardLink     *CardElementUrl `json:"card_link,omitempty"`
	Header       *CardHeader     `json:"header,omitempty"`
	Elements     []interface{}   `json:"elements,omitempty"`
	I18nElements *I18nElement    `json:"i18n_elements,omitempty"`
}

func (this *InteractiveMessage) GetType() string {
	return MessageType_Interactive
}

func (this *InteractiveMessage) GetCard() interface{} {
	return this
}

func NewSimpleInteractiveMessage(title string, content string) *InteractiveMessage {
	return &InteractiveMessage{
		Config: NewCardConfig(),
		Header: NewCardHeader(title),
		Elements: []interface{}{
			NewCardElementText(content),
		},
	}
}

func NewLinkInteractiveMessage(title, content, url string) *InteractiveMessage {
	return &InteractiveMessage{
		Config: NewCardConfig(),
		Header: NewCardHeader(title),
		Elements: []interface{}{
			NewCardElementText(content),
			NewCardElementHR(),
			map[string]interface{}{
				"tag":        "markdown",
				"content":    "** &#62; **",
				"text_align": "right",
				"text_size":  "small",
			},
		},
		CardLink: NewCardElementUrl(url),
	}
}

type CardConfig struct {
	WideScreenMode bool `json:"wide_screen_mode,omitempty"`
	EnableForward  bool `json:"enable_forward,omitempty"`
}

func NewCardConfig() *CardConfig {
	return &CardConfig{
		WideScreenMode: true, // 默认开启宽屏模式
	}
}

type CardElementUrl struct {
	Url        string `json:"url,omitempty"`
	AndroidUrl string `json:"android_url,omitempty"`
	IosUrl     string `json:"ios_url,omitempty"`
	PcUrl      string `json:"pc_url,omitempty"`
}

func NewCardElementUrl(url string) *CardElementUrl {
	return &CardElementUrl{
		Url: url,
	}
}

type CardHeader struct {
	Title    *CardHeaderTitle `json:"title,omitempty"`
	Template string           `json:"template,omitempty"`
}

func NewCardHeader(title string) *CardHeader {
	return &CardHeader{
		Title: &CardHeaderTitle{
			Tag:     "plain_text",
			Content: title,
		},
		Template: "blue", // 默认蓝色 https://open.feishu.cn/document/feishu-cards/card-components/content-components/title
	}
}

type CardHeaderTitle struct {
	Tag     string    `json:"tag,omitempty"` // 固定值 plain_text
	Content string    `json:"content,omitempty"`
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

type CardElementText struct {
	Tag  string                  `json:"tag,omitempty"`
	Text *CardElementTextContent `json:"text,omitempty"` // 配置卡片的普通文本信息。
	Icon *CardElementTextIcon    `json:"icon,omitempty"` // 添加图标作为文本前缀图标
}

func NewCardElementText(content string) *CardElementText {
	return &CardElementText{
		Tag: "div",
		Text: &CardElementTextContent{
			Tag:     "plain_text",
			Content: content,
		},
	}
}

type CardElementTextContent struct {
	Tag       string `json:"tag,omitempty"`        // 文本类型的标签 plain_text：普通文本内容或表情 lark_md：支持部分 Markdown 语法的文本内容。详情参考下文 lark_md 支持的 Markdown 语法
	Content   string `json:"content,omitempty"`    // 文本内容。当 tag 为 lark_md 时，支持部分 Markdown 语法的文本内容。详情参考下文 lark_md 支持的 Markdown 语法。
	TextSize  string `json:"text_size,omitempty"`  // 文本大小
	TextColor string `json:"text_color,omitempty"` // 文本的颜色。仅在 tag 为 plain_text 时生效。可取值： default：客户端浅色主题模式下为黑色；客户端深色主题模式下为白色
	TextAlign string `json:"text_align,omitempty"` // 文本对齐方式。可取值： left：左对齐 center：居中对齐 right：右对齐
	Lines     int    `json:"lines,omitempty"`      // 内容最大显示行数
}

type CardElementTextIcon struct {
	Tag    string `json:"tag,omitempty"`     // 图标类型的标签。可取值： standard_icon：使用图标库中的图标。 custom_icon：使用用自定义图片作为图标。
	Token  string `json:"token,omitempty"`   // 图标库中图标的 token。当 tag 为 standard_icon 时生效。枚举值参见图标库。
	Color  string `json:"color,omitempty"`   // 图标的颜色。支持设置线性和面性图标（即 token 末尾为 outlined 或 filled 的图标）的颜色。当 tag 为 standard_icon 时生效。枚举值参见颜色枚举值。
	ImgKey string `json:"img_key,omitempty"` // 自定义前缀图标的图片 key。当 tag 为 custom_icon 时生效。 图标 key 的获取方式：调用上传图片接口，上传用于发送消息的图片，并在返回值中获取图片的 image_key。
}

// CardElementHR 分隔线
type CardElementHR struct {
	Tag string `json:"tag"`
}

func NewCardElementHR() *CardElementHR {
	return &CardElementHR{
		Tag: "hr",
	}
}
