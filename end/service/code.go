package service

var (
	OK      = 00000 // 正常
	UNKNOWN = 00001 // 未知异常

	UserNotExit = 00001 // 用户不存在
)
var ServiceCode = map[int]string {
	OK:          "",
	UNKNOWN:     "未知异常",
	UserNotExit: "不存在该数据",
}