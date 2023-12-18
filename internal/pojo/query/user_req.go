package query

type UserProfileReq struct {
	RoomId   int64  `form:"room_id" json:"room_id" xml:"room_id"  binding:"required"`
	RoomUser string `form:"room_user" json:"room_user" xml:"room_user"  binding:"required"`
}

type UserProfileResponse struct {
	Amount       int64          `form:"amount" json:"amount" xml:"amount"  binding:"required"`
	UserProfiles []*UserProfile `form:"user_profiles" json:"user_profiles" xml:"user_profiles"  binding:"required"`
}
type UserProfile struct {
	RoomUser string `form:"room_user" json:"room_user" xml:"room_user"  binding:"required"`
	Amount   int64  `form:"amount" json:"amount" xml:"amount"  binding:"required"`
}

type UpdateUserNameReq struct {
	RoomId      int64  `form:"room_id" json:"room_id" xml:"room_id"  binding:"required"`
	RoomUserOld string `form:"room_user_old" json:"room_user_old" xml:"room_user_old"  binding:"required"`
	RoomUserNew string `form:"room_user_new" json:"room_user_new" xml:"room_user_new"  binding:"required"`
}
