package handler

var (
	OK       = 00000 // 正常
	UNKNOWN  = 00001 // 未知异常
	BizError = 00002 // 业务错误


	NOLOGIN = 10000 // 没有登录
	NOPERM = 10000 // 没有权限
)

var ApiCode  = map[int]string {
	OK:      "",
	UNKNOWN: "未知异常",
	NOLOGIN: "没登录",
	NOPERM: "没权限",
}