package query

type UserProfileReq struct {
	RoomId   int64  `form:"room_id" json:"room_id" xml:"room_id"  binding:"required"`
	RoomName string `form:"room_name" json:"room_name" xml:"room_name"  binding:"required"`
	RoomUser string `form:"room_user" json:"room_user" xml:"room_user"  binding:"required"`
}

type UserProfileResponse struct {
	Amount int64 `form:"amount" json:"amount" xml:"amount"  binding:"required"`
}
