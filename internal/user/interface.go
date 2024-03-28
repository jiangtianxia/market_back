package user

type IUserService interface {
	// 用户登录
	UserLogin(req *UserLoginReq) (*UserLoginResp, error)

	// 退出登录
	UserLogout(req *UserLogoutReq) error
}
