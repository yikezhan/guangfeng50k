package controller

import (
	"github.com/gin-gonic/gin"
	"guangfeng/internal/common"
	"guangfeng/internal/pojo/query"
	"net/http"
)

func Ping(c *gin.Context) {
	v := c.DefaultQuery("v", "test") //参数中的信息
	OKResponse(c, "pong: "+v)
}
func UserProfile(c *gin.Context) {
	var req query.UserProfileReq
	if err := c.ShouldBindJSON(&req); err != nil {
		FailResponse(c, http.StatusBadRequest, "fail", gin.H{"error": err.Error()})
		return
	}
	amount, userProfiles := srv.QueryUserProfile(req.RoomId, req.RoomUser)
	OKResponse(c, query.UserProfileResponse{
		Amount:       amount,
		UserProfiles: userProfiles,
	})
}
func GameDetail(c *gin.Context) {
	OKResponse(c, "敬请期待！")
}

func ModifyName(c *gin.Context) {
	var req query.UpdateUserNameReq
	if err := c.ShouldBindJSON(&req); err != nil {
		FailResponse(c, http.StatusBadRequest, "fail", gin.H{"error": err.Error()})
		return
	}
	if srv.ModifyName(req.RoomId, req.RoomUserOld, req.RoomUserNew) {
		OKResponse(c, true)
		return
	}
	FailResponseRCode(c, common.SystemError)
}
