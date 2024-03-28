package user

type UserLoginResp struct {
	Openid string `json:"openid" form:"openid"` // 微信openid
	Token  string `json:"token" form:"token"`   // 用户token
}
