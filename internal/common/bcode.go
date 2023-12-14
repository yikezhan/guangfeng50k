package common

type RCode struct {
	Code   int
	Msg    string
	Detail interface{}
}

var SystemError = &RCode{
	Code:   9999,
	Msg:    "异常，，，",
	Detail: nil,
}
var RoomUserError = &RCode{
	Code:   1000000,
	Msg:    "房间或用户不存在",
	Detail: nil,
}
var RoomExistError = &RCode{
	Code:   1000001,
	Msg:    "房间名已存在",
	Detail: nil,
}
var RoomInitError = &RCode{
	Code:   1000002,
	Msg:    "房间初始化失败",
	Detail: nil,
}
var PlayerNoAllSubmitError = &RCode{
	Code:   2000001,
	Msg:    "结算中，，，",
	Detail: nil,
}
var WinScoreError = &RCode{
	Code:   2000002,
	Msg:    "赢分冲突",
	Detail: nil,
}
var FullScoreError = &RCode{
	Code:   2000003,
	Msg:    "满分冲突",
	Detail: nil,
}
var SurroundError = &RCode{
	Code:   2000004,
	Msg:    "包围冲突",
	Detail: nil,
}
var KingPunishError = &RCode{
	Code:   2000005,
	Msg:    "罚王过多",
	Detail: nil,
}
var GameConfirmError = &RCode{
	Code:   2000006,
	Msg:    "还有人没确认战绩",
	Detail: nil,
}
var NextGameFail = &RCode{
	Code:   2000007,
	Msg:    "下一局尚未开始",
	Detail: nil,
}
