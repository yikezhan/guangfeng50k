package controller

import (
	"github.com/gin-gonic/gin"
	"guangfeng/internal/common"
	"guangfeng/internal/pojo/query"
	"net/http"
)

func NextGame(c *gin.Context) {
	var req query.NextGameReq
	if err := c.ShouldBindJSON(&req); err != nil {
		FailResponse(c, http.StatusBadRequest, "fail", gin.H{"error": err.Error()})
		return
	}
	number, code := srv.NextGame(req.RoomID, req.WxId, req.Number)
	if code != nil {
		FailResponseRCode(c, code)
		return
	}
	if number == req.Number+1 {
		OKResponse(c, &query.NextGameResponse{
			Number: number,
		})
	} else {
		FailResponseRCode(c, common.NextGameFail)
	}
}
func SubmitGameData(c *gin.Context) {
	var gameResult query.SubmitGameResultReq
	if err := c.ShouldBindJSON(&gameResult); err != nil {
		FailResponse(c, http.StatusBadRequest, "fail", gin.H{"error": err.Error()})
		return
	}
	res := srv.SubmitGameData(gameResult)
	OKResponse(c, res)
}
func ConfirmGameResult(c *gin.Context) {
	var gameResult query.ConfirmGameResultReq
	if err := c.ShouldBindJSON(&gameResult); err != nil {
		FailResponse(c, http.StatusBadRequest, "fail", gin.H{"error": err.Error()})
		return
	}
	if ok := srv.ConfirmGameResult(gameResult); !ok {
		FailResponseRCode(c, common.SystemError)
		return
	}
	OKResponse(c, true)
}

func GetGameResult(c *gin.Context) {
	var req query.GetGameResultReq
	if err := c.ShouldBindJSON(&req); err != nil {
		FailResponse(c, http.StatusBadRequest, "fail", gin.H{"error": err.Error()})
		return
	}

	gameResults, code := srv.QueryGameResult(req.RoomID, req.Number)
	if code != nil {
		FailResponseRCode(c, code)
		return
	}
	resp := &query.GetGameResultResponse{
		GameResultList: gameResults,
	}
	OKResponse(c, resp)
}
