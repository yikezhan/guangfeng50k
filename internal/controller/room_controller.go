package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"guangfeng/internal/common"
	"guangfeng/internal/pojo/query"
	"net/http"
)

func QueryRoom(c *gin.Context) {
	var room query.RoomQueryReq
	if err := c.ShouldBindJSON(&room); err != nil {
		FailResponse(c, http.StatusBadRequest, "fail", gin.H{"error": err.Error()})
		return
	}
	roomInfo := srv.QueryRoomById(room.RoomId)
	if roomInfo == nil {
		FailResponseRCode(c, common.RoomUserError)
		return
	}
	var roomRule *query.RoomRule
	json.Unmarshal([]byte(roomInfo.RuleJSON), &roomRule)
	res := &query.RoomQueryResponse{
		RoomId:    roomInfo.ID,
		RoomName:  roomInfo.RoomName,
		RoomRules: *roomRule,
		OwnerWxId: roomInfo.OwnerWxID,
		Number:    roomInfo.Number,
	}
	OKResponse(c, res)
}
func CreateRoom(c *gin.Context) {
	var room query.CreateRoomReq
	if err := c.ShouldBindJSON(&room); err != nil {
		FailResponse(c, http.StatusBadRequest, "fail", gin.H{"error": err.Error()})
		return
	}
	if ok := srv.InsertOrUpdateRoom(room); !ok {
		FailResponseRCode(c, common.RoomExistError)
		return
	}
	OKResponse(c, true)
}

func UpdateRoom(c *gin.Context) {
	var room query.CreateRoomReq
	if err := c.ShouldBindJSON(&room); err != nil {
		FailResponse(c, http.StatusBadRequest, "fail", gin.H{"error": err.Error()})
		return
	}
	if ok := srv.InsertOrUpdateRoom(room); !ok {
		FailResponseRCode(c, common.RoomExistError)
		return
	}
	OKResponse(c, true)
}
func EnterRoom(c *gin.Context) {
	var req query.EnterRoomReq
	if err := c.ShouldBindJSON(&req); err != nil {
		FailResponse(c, http.StatusBadRequest, "fail", gin.H{"error": err.Error()})
		return
	}
	room := srv.QueryRoom(req.RoomName)
	if room == nil {
		FailResponseRCode(c, common.RoomUserError)
		return
	}
	players := srv.QueryRoomUser(room.ID)
	for _, v := range players {
		if v.WxID == req.WxID {
			OKResponse(c, true)
			return
		}
	}
	if len(players) < 4 {
		if srv.CreateRoomUser(room.ID, req.WxID, req.WxImage, req.WxName) {
			OKResponse(c, true)
			return
		} else {
			FailResponseRCode(c, common.SystemError)
		}
	}
	FailResponseRCode(c, common.RoomFullError)
}
