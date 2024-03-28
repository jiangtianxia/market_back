package user

type UserLoginReq struct {
	Code string `json:"code" form:"code" binding:"required"` // 微信小程序code
}

type UserLogoutReq struct {
	Openid string `json:"openid" form:"openid" binding:"required"` // 微信openid
}
