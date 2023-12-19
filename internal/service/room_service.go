package service

import (
	"encoding/json"
	"guangfeng/internal/common"
	"guangfeng/internal/model"
	"guangfeng/internal/pojo/query"
	"time"
)

func (s *Service) QueryRoomById(roomId int64) *model.RoomTab {
	return s.dao.QueryRoomById(roomId)
}
func (s *Service) QueryRoom(roomName string) *model.RoomTab {
	return s.dao.QueryRoom(roomName)
}
func (s *Service) InsertOrUpdateRoom(req query.CreateRoomReq) bool {
	rules, _ := json.Marshal(req.RoomRules)
	room := &model.RoomTab{
		ID:         req.RoomId,
		RoomName:   req.RoomName,
		RuleJSON:   string(rules),
		OwnerWxID:  req.OwnerWxId,
		UpdateTime: time.Now().Unix(),
		IsDelete:   common.Valid,
	}
	if room.ID == 0 {
		return s.dao.CreateRoom(room)
	} else {
		return s.dao.UpdateRoom(room)
	}
}

func (s *Service) QueryRoomUser(roomId int64) []*model.RoomUserTab {
	return s.dao.QueryRoomUsers(roomId)
}
func (s *Service) CreateRoomUser(roomId int64, wxId string, wxImage string, wxName string) bool {
	roomUser := &model.RoomUserTab{
		RoomID:     roomId,
		WxID:       wxId,
		WxImage:    wxImage,
		WxUserName: wxName,
		Status:     1,
	}
	return s.dao.CreateRoomUser(roomUser)
}
