package query

type CreateRoomReq struct {
	RoomId    int64    `form:"room_id" json:"room_id" xml:"room_id"`
	RoomName  string   `form:"room_name" json:"room_name" xml:"room_name"  binding:"required"`
	RoomRules RoomRule `form:"room_rules" json:"room_rules" xml:"room_rules"  binding:"required"`
	OwnerWxId string   `form:"owner_wx_id" json:"owner_wx_id" xml:"owner_wx_id"  binding:"required"`
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
type RoomQueryReq struct {
	RoomId int64 `form:"room_id" json:"room_id" xml:"room_id"`
}
type RoomQueryResponse struct {
	RoomId    int64    `form:"room_id" json:"room_id" xml:"room_id"`
	RoomName  string   `form:"room_name" json:"room_name" xml:"room_name"  binding:"required"`
	RoomRules RoomRule `form:"room_rules" json:"room_rules" xml:"room_rules"  binding:"required"`
	OwnerWxId string   `form:"owner_wx_id" json:"owner_wx_id" xml:"owner_wx_id"  binding:"required"`
	Number    int64    `form:"number" json:"number" xml:"number"  binding:"required"`
}
type EnterRoomReq struct {
	RoomName string `form:"room_name" json:"room_name" xml:"room_name"  binding:"required"`
	WxID     string `form:"wx_id" json:"wx_id" xml:"wx_id"  binding:"required"`
	WxName   string `form:"wx_name" json:"wx_name" xml:"wx_name"  binding:"required"`
	WxImage  string `form:"wx_image" json:"wx_image" xml:"wx_image"`
}
type EnterRoomResponse struct {
	RoomId int64 `form:"room_id" json:"room_id" xml:"room_id"  binding:"required"`
	Number int64 `form:"number" json:"number" xml:"number"  binding:"required"`
}
