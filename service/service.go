package service

import "market_back/internal/user"

var (
	UserService user.IUserService
)

func init() {
	UserService = &userService{}
}
