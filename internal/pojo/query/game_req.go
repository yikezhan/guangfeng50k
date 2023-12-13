package query

type SubmitGameResultReq struct {
	ResultID   int64      `form:"result_id" json:"result_id" xml:"result_id"`
	RoomID     int64      `form:"room_id" json:"room_id" xml:"room_id"  binding:"required"`
	Number     int64      `form:"number" json:"number" xml:"number"  binding:"required"`
	RoomUser   string     `form:"room_user" json:"room_user" xml:"room_user"  binding:"required"`
	GameResult GameResult `form:"game_result" json:"game_result" xml:"game_result"  binding:"required"`
}
type GameResult struct {
	FiveBoom       int64 `form:"five_boom" json:"five_boom" xml:"five_boom"  binding:"required"`
	SixBoom        int64 `form:"six_boom" json:"six_boom" xml:"six_boom"  binding:"required"`
	SevenBoom      int64 `form:"seven_boom" json:"seven_boom" xml:"seven_boom"  binding:"required"`
	EightBoom      int64 `form:"eight_boom" json:"eight_boom" xml:"eight_boom"  binding:"required"`
	NineBoom       int64 `form:"nine_boom" json:"nine_boom" xml:"nine_boom"  binding:"required"`
	WinScore       int64 `form:"win_score" json:"win_score" xml:"win_score"  binding:"required"`
	FullScore      int64 `form:"full_score" json:"full_score" xml:"full_score"  binding:"required"`
	SurroundScore  int64 `form:"surround_score" json:"surround_score" xml:"surround_score"  binding:"required"`
	KingPunishment int64 `form:"king_punishment" json:"king_punishment" xml:"king_punishment"  binding:"required"`
}
type ConfirmGameResultReq struct {
	ResultID int64 `form:"result_id" json:"result_id" xml:"result_id"`
}
type CalGameResultReq struct {
	RoomID   int64  `form:"room_id" json:"room_id" xml:"room_id"  binding:"required"`
	Number   int64  `form:"number" json:"number" xml:"number"  binding:"required"`
	RoomUser string `form:"room_user" json:"room_user" xml:"room_user"  binding:"required"`
}
type CalGameResultResponse struct {
}
