package common

type RCode struct {
	Code   int
	Msg    string
	Detail interface{}
}

var RoomUserError = &RCode{
	Code:   1000000,
	Msg:    "房间或用户不存在",
	Detail: nil,
}
var SystemError = &RCode{
	Code:   1000001,
	Msg:    "异常，，，",
	Detail: nil,
}
var PlayerNoAllSubmitError = &RCode{
	Code:   7000001,
	Msg:    "结算中，，，",
	Detail: nil,
}
var WinScoreError = &RCode{
	Code:   7000002,
	Msg:    "赢分冲突",
	Detail: nil,
}
var FullScoreError = &RCode{
	Code:   7000003,
	Msg:    "满分冲突",
	Detail: nil,
}
var SurroundError = &RCode{
	Code:   7000004,
	Msg:    "包围冲突",
	Detail: nil,
}
var KingPunishError = &RCode{
	Code:   7000005,
	Msg:    "罚王过多",
	Detail: nil,
}
