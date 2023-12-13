package service

import (
	"encoding/json"
	"entryTask/internal/common"
	"entryTask/internal/model"
	"entryTask/internal/pojo/query"
	"entryTask/internal/util"
)

func (s *Service) InsertOrUpdateGameResult(req query.SubmitGameResultReq) bool {
	resultJson, _ := json.Marshal(req.GameResult)
	gameResult := &model.GameResultTab{
		ID:         req.ResultID,
		RoomID:     req.RoomID,
		Number:     req.Number,
		RoomUser:   req.RoomUser,
		ResultJSON: string(resultJson),
		Amount:     0,
	}
	if req.ResultID == 0 {
		util.ReflectBuildDefaultTimeAndValid(&gameResult)
		return s.dao.CreateGameResult(gameResult)
	} else {
		return s.dao.UpdateGameResult(gameResult)
	}
}
func (s *Service) ConfirmGameResult(req query.ConfirmGameResultReq) bool {
	return s.dao.UpdateStatus(req.ResultID)
}
func (s *Service) CalGameResult(req query.CalGameResultReq) (bool, *common.RCode) {
	res := s.dao.QueryGameResult(req.RoomID, req.Number)
	if ok, code := validResult(res); !ok {
		return false, code
	}
	room := s.dao.QueryRoomById(req.RoomID)
	if room == nil {
		return false, common.SystemError
	}
	roomRule := query.RoomRule{}
	if err := json.Unmarshal([]byte(room.RuleJSON), roomRule); err != nil {
		return false, common.SystemError
	}
	amountMap := make(map[int64]int64)
	var kingPunishAmountTotal int64
	for _, v := range res {
		result := &query.GameResult{}
		if err := json.Unmarshal([]byte(v.ResultJSON), result); err != nil {
			return false, common.SystemError
		}
		var kingPunishAmount int64
		amountMap[v.ID], kingPunishAmount = CalAmount(roomRule, result)
		kingPunishAmountTotal = kingPunishAmountTotal + kingPunishAmount
	}
	for id, amount := range amountMap {
		amountMap[id] = amount - kingPunishAmountTotal/4
		if ok := s.dao.UpdateAmount(id, amount); !ok {
			return false, common.SystemError
		}
	}
	return true, nil
}

func CalAmount(rule query.RoomRule, result *query.GameResult) (int64, int64) {
	var amount int64
	amount = result.FiveBoom*rule.FiveBoom +
		result.SixBoom*rule.SixBoom +
		result.SevenBoom*rule.SevenBoom +
		result.EightBoom*rule.EightBoom +
		result.NineBoom*rule.NineBoom
	if result.WinScore == common.Yes {
		amount = amount + rule.WinScore
	}
	if result.WinScore == common.No {
		amount = amount - rule.WinScore
	}
	if result.FullScore == common.Yes {
		amount = amount + rule.FullScore
	}
	if result.FullScore == common.No {
		amount = amount - rule.FullScore
	}
	if result.SurroundScore == common.Yes {
		amount = amount + rule.SurroundScore
	}
	if result.SurroundScore == common.No {
		amount = amount - rule.SurroundScore
	}
	// 先全罚，再拿回自己的部分
	kingPunishAmount := result.KingPunishment * result.KingPunishment * common.PlayerNumber
	amount = amount - kingPunishAmount
	return amount, kingPunishAmount
}

func validResult(res []model.GameResultTab) (bool, *common.RCode) {
	if len(res) != common.PlayerNumber {
		return false, common.PlayerNoAllSubmitError
	}
	var winScoreCount int64
	var FullScoreCount int64
	var SurroundScoreCount int64
	var KingPunishment int64
	for _, v := range res {
		gameResult := &query.GameResult{}
		if err := json.Unmarshal([]byte(v.ResultJSON), gameResult); err != nil {
			return false, common.SystemError
		}
		if gameResult.WinScore == common.Yes {
			winScoreCount++
		}
		if gameResult.WinScore == common.No {
			winScoreCount--
		}
		if gameResult.FullScore == common.Yes {
			FullScoreCount++
		}
		if gameResult.FullScore == common.No {
			FullScoreCount--
		}
		if gameResult.SurroundScore == common.Yes {
			SurroundScoreCount++
		}
		if gameResult.SurroundScore == common.No {
			SurroundScoreCount--
		}
		KingPunishment = KingPunishment + gameResult.KingPunishment
	}
	if winScoreCount != 0 {
		return false, common.WinScoreError
	}
	if FullScoreCount != 0 {
		return false, common.FullScoreError
	}
	if SurroundScoreCount != 0 {
		return false, common.SurroundError
	}
	if KingPunishment > 4 {
		return false, common.KingPunishError
	}
	return true, nil
}
