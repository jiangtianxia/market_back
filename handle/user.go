package handle

import (
	"market_back/internal/user"
	mlog "market_back/logger"
	"market_back/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandle struct{}

// UserLogin
//
//	@tags			用户相关
//	@Summary		用户登录
//	@Description	用户登录
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Param			request	body		user.UserLoginReq	true	"body"
//	@Success		200		{object}	user.UserLoginResp
//	@Router			/user/login [post]
func (*userHandle) UserLogin(ctx *gin.Context) {
	var req user.UserLoginReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, ErrParam)
		return
	}

	resp, err := service.UserService.UserLogin(&req)
	if err != nil {
		mlog.Errorf("[UserLogin] user login failed, err: %v", err)
		ctx.AbortWithStatusJSON(http.StatusOK, Result(err))
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, Result(resp))
}

// UserLogout
//
//	@tags			用户相关
//	@Summary		退出登录
//	@Description	退出登录
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Param			Authorization	header		string				true	"Authorization"
//	@Param			request			body		user.UserLogoutReq	true	"body"
//	@Success		200				string		"success"
//	@Router			/user/logout [post]
func (*userHandle) UserLogout(ctx *gin.Context) {
	panic("unimplemented")
}
