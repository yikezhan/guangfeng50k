package controller

import (
	"entryTask/internal/common"
	"entryTask/internal/pojo/query"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRoom(c *gin.Context) {
	var room query.RoomReq
	if err := c.ShouldBindJSON(&room); err != nil {
		FailResponse(c, http.StatusBadRequest, "fail", gin.H{"error": err.Error()})
		return
	}
	res := srv.InsertOrUpdateRoom(room)
	OKResponse(c, res)
}
func UpdateRoom(c *gin.Context) {
	var room query.RoomReq
	if err := c.ShouldBindJSON(&room); err != nil {
		FailResponse(c, http.StatusBadRequest, "fail", gin.H{"error": err.Error()})
		return
	}
	res := srv.InsertOrUpdateRoom(room)
	OKResponse(c, res)
}
func EnterRoom(c *gin.Context) {
	var enterRoom query.EnterRoomReq
	if err := c.ShouldBindJSON(&enterRoom); err != nil {
		FailResponse(c, http.StatusBadRequest, "fail", gin.H{"error": err.Error()})
		return
	}
	res := srv.QueryRoom(enterRoom.RoomName)
	if res.RoomOwner == enterRoom.RoomUser ||
		res.RoomUser1 == enterRoom.RoomUser ||
		res.RoomUser2 == enterRoom.RoomUser ||
		res.RoomUser3 == enterRoom.RoomUser {
		OKResponse(c, res)
		return
	}
	FailResponseRCode(c, common.RoomUserError)
}
