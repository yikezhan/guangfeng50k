package controller

import (
	"github.com/gin-gonic/gin"
	"guangfeng/internal/common"
	"guangfeng/internal/pojo/query"
	"net/http"
)

func CreateRoom(c *gin.Context) {
	var room query.RoomReq
	if err := c.ShouldBindJSON(&room); err != nil {
		FailResponse(c, http.StatusBadRequest, "fail", gin.H{"error": err.Error()})
		return
	}
	if ok := srv.InsertOrUpdateRoom(room); !ok {
		FailResponseRCode(c, common.RoomExistError)
		return
	}
	if ok := srv.InitGameData(room.RoomName, 1); !ok {
		FailResponseRCode(c, common.RoomInitError)
		return
	}
	OKResponse(c, true)
}

func UpdateRoom(c *gin.Context) {
	var room query.RoomReq
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
	if room.RoomOwner == req.RoomUser ||
		room.RoomUser1 == req.RoomUser ||
		room.RoomUser2 == req.RoomUser ||
		room.RoomUser3 == req.RoomUser {
		res, code := srv.LatestGameResult(room.ID, req.RoomUser)
		if code != nil {
			FailResponseRCode(c, code)
			return
		}
		OKResponse(c, &query.EnterRoomResponse{
			ID:       res.ID,
			RoomId:   res.RoomID,
			RoomUser: res.RoomUser,
			Number:   res.Number,
		})
		return
	}
	FailResponseRCode(c, common.RoomUserError)
}
