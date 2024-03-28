package handle

import "github.com/gin-gonic/gin"

var (
	UserHandle IUserHandle
)

func init() {
	UserHandle = &userHandle{}
}

// 用户相关
type IUserHandle interface {
	// 用户登录
	UserLogin(ctx *gin.Context)

	// 退出登录
	UserLogout(ctx *gin.Context)
}
