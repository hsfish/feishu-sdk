package feishu_sdk

const (
	api_Chats_Query  = "https://open.feishu.cn/open-apis/im/v1/chats"
	api_Chats_Search = "https://open.feishu.cn/open-apis/im/v1/chats/search"
)

type SearchChatResult struct {
	List      []*Chat `json:"items"`
	PageToken string  `json:"page_token"`
	HasMore   bool    `json:"has_more"`
}

type Chat struct {
	ChatId      string `json:"chat_id"`
	Avatar      string `json:"avatar"`
	Name        string `json:"name"`
	Description string `json:"description"`
	OwnerId     string `json:"owner_id"`
	OwnerIdType string `json:"owner_id_type"`
	External    bool   `json:"external"`
	TenantKey   string `json:"tenant_key"`
}

// 获取用户或机器人所在的群列表
func (this *Sdk) QueryChat(userIdType, pageToken string, pageSize int, all ...bool) (*SearchChatResult, error) {
	result := &SearchChatResult{}
	resp := &baseResultWithData{Data: &SearchChatResult{}}

	rawQuery := map[string]interface{}{}
	if userIdType != "" {
		rawQuery["user_id_type"] = userIdType
	}
	if pageToken != "" {
		rawQuery["page_token"] = pageToken
	}
	if pageSize <= 0 {
		pageSize = 100
	}
	rawQuery["page_size"] = pageSize

	isAll := append(all, false)[0]
	for {
		if _, err := this.GetWithAuth(api_Chats_Query, rawQuery, nil, resp); err != nil {
			return result, err
		}

		r := resp.Data.(*SearchChatResult)
		result.HasMore = r.HasMore
		result.PageToken = r.PageToken
		if len(r.List) > 0 {
			result.List = append(result.List, r.List...)
		}
		if !result.HasMore || !isAll {
			return result, nil
		}
		rawQuery["page_token"] = result.PageToken
	}
}

// 搜索对用户或机器人可见的群列表
func (this *Sdk) SearchChat(userIdType, query, pageToken string, pageSize int, all ...bool) (*SearchChatResult, error) {
	result := &SearchChatResult{}
	resp := &baseResultWithData{Data: &SearchChatResult{}}

	rawQuery := map[string]interface{}{}
	if userIdType != "" {
		rawQuery["user_id_type"] = userIdType
	}
	if pageToken != "" {
		rawQuery["page_token"] = pageToken
	}
	if query != "" {
		rawQuery["query"] = query
	}
	if pageSize <= 0 {
		pageSize = 100
	}
	rawQuery["page_size"] = pageSize

	isAll := append(all, false)[0]
	for {
		if _, err := this.GetWithAuth(api_Chats_Search, rawQuery, nil, resp); err != nil {
			return result, err
		}
		r := resp.Data.(*SearchChatResult)
		result.HasMore = r.HasMore
		result.PageToken = r.PageToken
		if len(r.List) > 0 {
			result.List = append(result.List, r.List...)
		}
		if !result.HasMore || !isAll {
			return result, nil
		}
		rawQuery["page_token"] = result.PageToken
	}
}
