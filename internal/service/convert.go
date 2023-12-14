package service

import (
	"encoding/json"
	"entryTask/internal/model"
	"entryTask/internal/pojo/query"
)

func convertGameResultTabToUserGameResult(data []model.GameResultTab) []query.UserGameResult {
	res := make([]query.UserGameResult, 0)
	for _, v := range data {
		var gameResult query.GameResult
		json.Unmarshal([]byte(v.ResultJSON), gameResult)

		tmp := query.UserGameResult{
			ResultID:   v.ID,
			RoomID:     v.RoomID,
			Number:     v.Number,
			RoomUser:   v.RoomUser,
			GameResult: gameResult,
		}
		res = append(res, tmp)
	}
	return res
}
