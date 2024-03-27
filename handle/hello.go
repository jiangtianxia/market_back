package handle

import "github.com/gin-gonic/gin"

// Hello
//
//	@tags			hello
//	@Summary		hello
//	@Description	hello
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Success		200	{string}	string	"success"
//	@Router			/hello [get]
func Hello(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "hello",
	})
}
