package service

import (
	"encoding/json"
	"guangfeng/internal/common"
	"guangfeng/internal/model"
	"guangfeng/internal/pojo/query"
	"time"
)

func (s *Service) QueryUserProfile(roomID int64, roomUser string) (int64, []*query.UserProfile) {
	gameResults := s.dao.AllGameResult(roomID)
	userProfiles := make([]*query.UserProfile, 0)
	userMap := make(map[string]int64, 0)
	for _, v := range gameResults {
		userMap[v.RoomUserID] = userMap[v.RoomUserID] + v.Amount
	}
	for k, v := range userMap {
		userProfiles = append(userProfiles, &query.UserProfile{
			RoomUser: k,
			Amount:   v,
		})
	}
	return userMap[roomUser], userProfiles
}
func (s *Service) QueryGameResult(roomID int64, number int64) ([]query.UserGameResult, *common.RCode) {

	res := s.dao.QueryGameResult(roomID, number)
	if ok, code := validResult(res); !ok {
		return nil, code
	}
	return convertGameResultTabToUserGameResult(res), nil
}

func (s *Service) LatestGameResult(roomId int64, roomUser string) (*model.GameResultTab, *common.RCode) {
	res := s.dao.LatestGameResult(roomId, roomUser)
	if len(res) != 1 {
		return nil, common.SystemError
	}
	return res[0], nil
}
func (s *Service) NextGame(roomId int64, wxId string) (int64, *common.RCode) {
	room := s.dao.QueryRoomById(roomId)
	latestGame := s.dao.LatestGameResult(roomId, wxId)[0]
	if room.OwnerWxID == wxId {
		latestGameResult := s.dao.QueryGameResult(roomId, latestGame.Number)
		for _, v := range latestGameResult {
			if v.Status != common.Confirm {
				return latestGame.Number, common.GameConfirmError
			}
		}
		return latestGame.Number + 1, nil
	}
	return latestGame.Number, nil
}
func (s *Service) SubmitGameData(req query.SubmitGameResultReq) bool {
	resultJson, _ := json.Marshal(req.GameResult)
	gameResult := &model.GameResultTab{
		ID:         req.ResultID,
		ResultJSON: string(resultJson),
		UpdateTime: time.Now().Unix(),
		Status:     common.Cal,
	}
	result := s.dao.QueryGameResult(req.RoomID, req.Number)
	for _, v := range result {
		if v.RoomUserID == req.WxId {
			gameResult.ID = v.ID
			break
		}
	}
	if gameResult.ID == 0 {
		if !s.dao.UpdateGameResult(gameResult) {
			return false
		}
	} else {
		if !s.dao.CreateGameResult(gameResult) {
			return false
		}
	}
	result = s.dao.QueryGameResult(req.RoomID, req.Number)
	if len(result) == 4 {
		s.CalGameResult(req.RoomID, req.Number)
	}
	return true
}
func (s *Service) ConfirmGameResult(req query.ConfirmGameResultReq) bool {
	return s.dao.ConfirmResult(req.ResultID, req.Amount)
}
func (s *Service) CalGameResult(roomID int64, number int64) (bool, *common.RCode) {
	res := s.dao.QueryGameResult(roomID, number)
	room := s.dao.QueryRoomById(roomID)
	if room == nil {
		return false, common.SystemError
	}
	var roomRule *query.RoomRule
	if err := json.Unmarshal([]byte(room.RuleJSON), &roomRule); err != nil {
		return false, common.SystemError
	}
	amountMap := make(map[int64]int64)
	for _, v := range res {
		amountMap[v.ID] = 0
	}
	for _, v := range res {
		if ok, code := CalAmount(roomRule, v, amountMap); !ok {
			return false, code
		}
	}
	for id, amount := range amountMap {
		s.dao.UpdateAmount(id, amount)
	}
	return true, nil
}

func CalAmount(rule *query.RoomRule, v model.GameResultTab, amountMap map[int64]int64) (bool, *common.RCode) {

	result := &query.GameResult{}
	if err := json.Unmarshal([]byte(v.ResultJSON), result); err != nil {
		return false, common.SystemError
	}
	boomAmount := result.FiveBoom*rule.FiveBoom +
		result.SixBoom*rule.SixBoom +
		result.SevenBoom*rule.SevenBoom +
		result.EightBoom*rule.EightBoom +
		result.NineBoom*rule.NineBoom
	// 先全罚，再拿回自己的部分
	kingPunishAmount := result.KingPunishment * rule.KingScore
	for id, _ := range amountMap {
		if id == v.ID {
			amountMap[id] = amountMap[id] + boomAmount*(common.PlayerNumber-1) - kingPunishAmount*(common.PlayerNumber-1)
		} else {
			amountMap[id] = amountMap[id] - boomAmount + kingPunishAmount
		}
	}
	if result.WinScore == common.Win {
		amountMap[v.ID] = amountMap[v.ID] + rule.WinScore
	}
	if result.WinScore == common.Fail {
		amountMap[v.ID] = amountMap[v.ID] - rule.WinScore
	}
	if result.FullScore == common.Win {
		amountMap[v.ID] = amountMap[v.ID] + rule.FullScore
	}
	if result.FullScore == common.Fail {
		amountMap[v.ID] = amountMap[v.ID] - rule.FullScore
	}
	if result.SurroundScore == common.Win {
		amountMap[v.ID] = amountMap[v.ID] + rule.SurroundScore
	}
	if result.SurroundScore == common.Fail {
		amountMap[v.ID] = amountMap[v.ID] - rule.SurroundScore
	}
	return true, nil
}

func validResult(res []model.GameResultTab) (bool, *common.RCode) {
	var winScoreCount int64
	var FullScoreCount int64
	var SurroundScoreCount int64
	var KingPunishment int64
	for _, v := range res {
		if v.Status == common.Draft {
			return false, common.PlayerNoAllSubmitError
		}
		gameResult := &query.GameResult{}
		if err := json.Unmarshal([]byte(v.ResultJSON), gameResult); err != nil {
			return false, common.SystemError
		}
		if gameResult.WinScore == common.Win {
			winScoreCount++
		}
		if gameResult.WinScore == common.Fail {
			winScoreCount--
		}
		if gameResult.FullScore == common.Win {
			FullScoreCount++
		}
		if gameResult.FullScore == common.Fail {
			FullScoreCount--
		}
		if gameResult.SurroundScore == common.Win {
			SurroundScoreCount++
		}
		if gameResult.SurroundScore == common.Fail {
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
