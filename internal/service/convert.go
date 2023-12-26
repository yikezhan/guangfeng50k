package service

import (
	"encoding/json"
	"guangfeng/internal/model"
	"guangfeng/internal/pojo/query"
)

func convertGameResultTabToUserGameResultList(data []model.GameResultTab) []*query.UserGameResult {
	res := make([]*query.UserGameResult, 0)
	for _, v := range data {
		tmp := convertGameResultTabToUserGameResult(&v)
		res = append(res, tmp)
	}
	return res
}

func convertGameResultTabToUserGameResult(v *model.GameResultTab) *query.UserGameResult {
	var gameResult *query.GameResult
	json.Unmarshal([]byte(v.ResultJSON), &gameResult)

	tmp := &query.UserGameResult{
		ResultID:   v.ID,
		RoomID:     v.RoomID,
		Number:     v.Number,
		WxID:       v.WxID,
		GameResult: *gameResult,
		Amount:     v.Amount,
	}
	return tmp
}
