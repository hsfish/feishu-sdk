package feishu_sdk

const (
	api_User_Batch_Get_Id_V3 = "https://open.feishu.cn/open-apis/contact/v3/users/batch_get_id"
)

type UserIdType string

// 用户类型
const (
	UserIdType_Open_Id  = UserIdType("open_id")  //  用户的 open id
	UserIdType_Union_Id = UserIdType("union_id") // 用户的 union id
	UserIdType_UserId   = UserIdType("user_id")  // 用户的 user id
)

type BatchGetUserIdResponse struct {
	UserList []UserIdDetail `json:"user_list"`
}

type UserIdDetail struct {
	UserId string `json:"user_id"` //用户id，值为user_id_type所指定的类型。如果查询的手机号、邮箱不存在，或者无权限查看对应的用户，则此项为空。
	Mobile string `json:"mobile"`  //手机号，通过手机号查询时返回
	Email  string `json:"email"`   //邮箱，通过邮箱查询时返回
}

func (this *sdk) BatchGetUserId(emails []string, mobiles []string, idType ...UserIdType) (*BatchGetUserIdResponse, error) {
	resp := &baseResultWithData{Data: &BatchGetUserIdResponse{}}

	if _, err := this.PostWithAuth(api_User_Batch_Get_Id_V3, map[string]interface{}{
		"user_id_type": append(idType, UserIdType_UserId)[0],
	}, map[string]interface{}{
		"emails":  emails,
		"mobiles": mobiles,
	}, resp); err != nil {
		return nil, err
	}

	return resp.Data.(*BatchGetUserIdResponse), nil
}
