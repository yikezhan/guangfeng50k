package service

import (
	"encoding/json"
	"entryTask/internal/common"
	"entryTask/internal/model"
	"entryTask/internal/pojo/query"
	"strconv"
	"time"
)

func (s *Service) QueryRoom(roomName string) *model.RoomTab {
	return s.dao.QueryRoom(roomName)
}
func (s *Service) InsertOrUpdateRoom(req query.RoomReq) bool {
	rules, _ := json.Marshal(req.RoomRules)
	room := &model.RoomTab{
		ID:         req.RoomId,
		RoomName:   req.RoomName,
		Password:   req.Password,
		RuleJSON:   string(rules),
		RoomOwner:  "房主",
		RoomUser1:  defaultName(req.RoomUser1, 1),
		RoomUser2:  defaultName(req.RoomUser2, 2),
		RoomUser3:  defaultName(req.RoomUser3, 3),
		UpdateTime: time.Now().Unix(),
		IsDelete:   common.Valid,
	}
	if room.ID == 0 {
		return s.dao.CreateRoom(room)
	} else {
		return s.dao.UpdateRoom(room)
	}
}

func defaultName(name string, num int) string {
	if name != "" {
		return name
	}
	// 嘿嘿，这默认名nice吧
	switch num {
	case 1:
		return "花开富贵"
	case 2:
		return "上善若水"
	case 3:
		return "美好明天"
	default:
		return "皮卡丘" + strconv.Itoa(1) + "号"
	}
}
