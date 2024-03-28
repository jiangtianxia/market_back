package handle

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	// Code reply code
	Code = "code"
	// Message reply message
	Message = "message"
	// Data ...
	Data = "data"
	// RealError
	RealError = "real_error"
)

const (
	success = iota
	failed  = iota
)

var (
	// Success ...
	Success = gin.H{Code: failed, Message: "success"}

	// ErrParam  param error
	ErrParam = gin.H{Code: failed, Message: "param error"}
)

type MError struct {
	Code    int
	Message string
	Data    interface{}
}

// Result
func Result(result interface{}) gin.H {
	switch result.(type) {
	case *MError:
		err := result.(*MError)
		return gin.H{
			Code:    err.Code,
			Message: err.Message,
			Data:    err.Data,
		}
	case MError:
		err := result.(MError)
		return gin.H{
			Code:    err.Code,
			Message: err.Message,
			Data:    err.Data,
		}
	case error:
		return gin.H{
			Code:      http.StatusInternalServerError,
			Message:   "内部错误, 请重试或联系管理员",
			RealError: result.(error).Error(),
		}
	default:
		return gin.H{
			Code:    success,
			Message: "success",
			Data:    result,
		}
	}
}
