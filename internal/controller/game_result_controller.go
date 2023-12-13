package controller

import (
	"entryTask/internal/pojo/query"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateGameResult(c *gin.Context) {
	var gameResult query.SubmitGameResultReq
	if err := c.ShouldBindJSON(&gameResult); err != nil {
		FailResponse(c, http.StatusBadRequest, "fail", gin.H{"error": err.Error()})
		return
	}
	res := srv.InsertOrUpdateGameResult(gameResult)
	OKResponse(c, res)
}
func UpdateGameResult(c *gin.Context) {
	var gameResult query.SubmitGameResultReq
	if err := c.ShouldBindJSON(&gameResult); err != nil {
		FailResponse(c, http.StatusBadRequest, "fail", gin.H{"error": err.Error()})
		return
	}
	res := srv.InsertOrUpdateGameResult(gameResult)
	OKResponse(c, res)
}
func ConfirmGameResult(c *gin.Context) {
	var gameResult query.ConfirmGameResultReq
	if err := c.ShouldBindJSON(&gameResult); err != nil {
		FailResponse(c, http.StatusBadRequest, "fail", gin.H{"error": err.Error()})
		return
	}
	res := srv.ConfirmGameResult(gameResult)
	OKResponse(c, res)
}

func Calculate(c *gin.Context) {
	var calReq query.CalGameResultReq
	if err := c.ShouldBindJSON(&calReq); err != nil {
		FailResponse(c, http.StatusBadRequest, "fail", gin.H{"error": err.Error()})
		return
	}
	var resp query.CalGameResultResponse
	ok, code := srv.CalGameResult(calReq)
	if !ok {
		FailResponseRCode(c, code)
		return
	}
	OKResponse(c, resp)
}
