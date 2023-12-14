package controller

import (
	"github.com/gin-gonic/gin"
	"guangfeng/internal/common"
	"guangfeng/internal/service"
	"net/http"
)

var (
	srv *service.Service
)

func Init() {
	srv = service.New()
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func OKResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "success",
		Data:    data,
	})
}
func FailResponse(c *gin.Context, code int, message string, data interface{}) {

	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func FailResponseRCode(c *gin.Context, rCode *common.RCode) {
	c.JSON(http.StatusOK, Response{
		Code:    rCode.Code,
		Message: rCode.Msg,
		Data:    rCode.Detail,
	})
}
