package service

import (
	"fmt"
	"market_back/internal/user"
	"market_back/store"
	"market_back/utils"
	"time"
)

type userService struct {
}

// UserLogin 用户登录
func (*userService) UserLogin(req *user.UserLoginReq) (*user.UserLoginResp, error) {
	// 获取微信openid
	openid, err := utils.GetOpenid(req.Code)
	if err != nil {
		return nil, fmt.Errorf("get openid failed: %v", err)
	}

	token := ""
	// 判断该openid对应的token是否存在于redis中
	if !store.ExistsKey(openid) {
		// 如果不存在，则生成新的token
		token, err = utils.GenerateToken(openid)
		if err != nil {
			return nil, fmt.Errorf("generate token failed, openid: %s, err: %v", openid, err)
		}
	} else {
		// 如果存在，则获取该openid对应的token
		token, err = store.GetKey(openid)
		if err != nil {
			return nil, fmt.Errorf("get token from redis failed, openid: %s, err: %v", openid, err)
		}
	}

	// 更新redis
	if err = store.SetKeyWithExpire(openid, token, time.Hour*24); err != nil {
		return nil, fmt.Errorf("set token to redis failed, openid: %s, token: %s, err: %v", openid, token, err)
	}

	return &user.UserLoginResp{
		Openid: openid,
		Token:  token,
	}, nil
}

// UserLogout implements user.IUserService.
func (*userService) UserLogout(req *user.UserLogoutReq) error {
	panic("unimplemented")
}
