package query

type SubmitGameResultReq struct {
	RoomID     int64      `form:"room_id" json:"room_id" xml:"room_id"  binding:"required"`
	Number     int64      `form:"number" json:"number" xml:"number"  binding:"required"`
	WxId       string     `form:"wx_id" json:"wx_id" xml:"wx_id"  binding:"required"`
	GameResult GameResult `form:"game_result" json:"game_result" xml:"game_result"`
}
type GameResult struct {
	FiveBoom       int64 `form:"five_boom" json:"five_boom" xml:"five_boom"`
	SixBoom        int64 `form:"six_boom" json:"six_boom" xml:"six_boom"`
	SevenBoom      int64 `form:"seven_boom" json:"seven_boom" xml:"seven_boom"`
	EightBoom      int64 `form:"eight_boom" json:"eight_boom" xml:"eight_boom"`
	NineBoom       int64 `form:"nine_boom" json:"nine_boom" xml:"nine_boom"`
	WinScore       int64 `form:"win_score" json:"win_score" xml:"win_score"`
	FullScore      int64 `form:"full_score" json:"full_score" xml:"full_score"`
	SurroundScore  int64 `form:"surround_score" json:"surround_score" xml:"surround_score"`
	KingPunishment int64 `form:"king_punishment" json:"king_punishment" xml:"king_punishment"`
}
type ConfirmGameResultReq struct {
	ResultID int64 `form:"result_id" json:"result_id" xml:"result_id"`
	Amount   int64 `form:"amount" json:"amount" xml:"amount"`
}
type CalGameResultReq struct {
	RoomID   int64  `form:"room_id" json:"room_id" xml:"room_id"  binding:"required"`
	Number   int64  `form:"number" json:"number" xml:"number"  binding:"required"`
	RoomUser string `form:"room_user" json:"room_user" xml:"room_user"  binding:"required"`
}

type NextGameReq struct {
	RoomID int64  `form:"room_id" json:"room_id" xml:"room_id"  binding:"required"`
	WxId   string `form:"wx_id" json:"wx_id" xml:"wx_id"  binding:"required"`
	Number int64  `form:"number" json:"number" xml:"number"  binding:"required"`
}
type NextGameResponse struct {
	Number int64 `form:"number" json:"number" xml:"number"  binding:"required"`
}

type GetGameResultReq struct {
	RoomID int64 `form:"room_id" json:"room_id" xml:"room_id"  binding:"required"`
	Number int64 `form:"number" json:"number" xml:"number"  binding:"required"`
}

type UserGameResult struct {
	ResultID   int64      `form:"result_id" json:"result_id" xml:"result_id"`
	RoomID     int64      `form:"room_id" json:"room_id" xml:"room_id"  binding:"required"`
	Number     int64      `form:"number" json:"number" xml:"number"  binding:"required"`
	WxID       string     `form:"wx_id" json:"wx_id" xml:"wx_id"  binding:"required"`
	GameResult GameResult `form:"game_result" json:"game_result" xml:"game_result"  binding:"required"`
	Amount     int64      `form:"amount" json:"amount" xml:"amount"`
}
type GetGameResultResponse struct {
	GameResultList []UserGameResult `form:"game_result_list" json:"game_result_list" xml:"game_result_list"  binding:"required"`
}
