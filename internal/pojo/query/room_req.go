package query

type RoomReq struct {
	RoomId    int64    `form:"room_id" json:"room_id" xml:"room_id"`
	RoomName  string   `form:"room_name" json:"room_name" xml:"room_name"  binding:"required"`
	Password  string   `form:"password" json:"password" xml:"password"`
	RoomRules RoomRule `form:"room_rules" json:"room_rules" xml:"room_rules"  binding:"required"`
	RoomUser1 string   `form:"room_user1" json:"room_user1" xml:"room_user1"  binding:"required"`
	RoomUser2 string   `form:"room_user2" json:"room_user2" xml:"room_user2"  binding:"required"`
	RoomUser3 string   `form:"room_user3" json:"room_user3" xml:"room_user3"  binding:"required"`
}
type RoomRule struct {
	FiveBoom      int64 `form:"five_boom" json:"five_boom" xml:"five_boom"  binding:"required"`
	SixBoom       int64 `form:"six_boom" json:"six_boom" xml:"six_boom"  binding:"required"`
	SevenBoom     int64 `form:"seven_boom" json:"seven_boom" xml:"seven_boom"  binding:"required"`
	EightBoom     int64 `form:"eight_boom" json:"eight_boom" xml:"eight_boom"  binding:"required"`
	NineBoom      int64 `form:"nine_boom" json:"nine_boom" xml:"nine_boom"  binding:"required"`
	WinScore      int64 `form:"win_score" json:"win_score" xml:"win_score"  binding:"required"`
	FullScore     int64 `form:"full_score" json:"full_score" xml:"full_score"  binding:"required"`
	SurroundScore int64 `form:"surround_score" json:"surround_score" xml:"surround_score"  binding:"required"`
	KingScore     int64 `form:"king_score" json:"king_score" xml:"king_score"  binding:"required"`
}
type EnterRoomReq struct {
	RoomName string `form:"room_name" json:"room_name" xml:"room_name"  binding:"required"`
	RoomUser string `form:"room_user" json:"room_user" xml:"room_user"  binding:"required"`
}
type EnterRoomResponse struct {
	ID       int64  `form:"id" json:"id" xml:"id"  binding:"required"`
	RoomId   int64  `form:"room_id" json:"room_id" xml:"room_id"  binding:"required"`
	RoomUser string `form:"room_user" json:"room_user" xml:"room_user"  binding:"required"`
}
