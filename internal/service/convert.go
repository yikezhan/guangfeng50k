package service

import (
	"encoding/json"
	"guangfeng/internal/model"
	"guangfeng/internal/pojo/query"
)

func convertGameResultTabToUserGameResultList(data []model.GameResultTab, roomUser []*model.RoomUserTab) []*query.UserGameResult {
	res := make([]*query.UserGameResult, 0)
	for _, v := range data {
		tmp := convertGameResultTabToUserGameResult(&v, roomUser)
		res = append(res, tmp)
	}
	return res
}

func convertGameResultTabToUserGameResult(v *model.GameResultTab, roomUsers []*model.RoomUserTab) *query.UserGameResult {
	var gameResult *query.GameResult
	json.Unmarshal([]byte(v.ResultJSON), &gameResult)
	tmp := &query.UserGameResult{
		ResultID:   v.ID,
		RoomID:     v.RoomID,
		Number:     v.Number,
		WxID:       v.WxID,
		GameResult: *gameResult,
		Status:     v.Status,
		Amount:     v.Amount,
	}
	for _, user := range roomUsers {
		if user.WxID == tmp.WxID {
			tmp.WxName = user.WxUserName
		}
	}
	return tmp
}
