package middlewares

import (
	"market_back/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CurrentLimit 限流拦截器
func CurrentLimit() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 判断令牌桶中是否有令牌
		if !utils.Bucket.Allow() {
			// 不允许访问
			ctx.Abort()
			ctx.JSON(http.StatusOK, gin.H{
				"code":    -1,
				"message": "服务器繁忙，请稍后重试",
			})
			return
		}
		ctx.Next()
	}
}
