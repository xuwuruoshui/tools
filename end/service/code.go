package service

// @SuffixName   ServiceCode
// @Description  接口相关

var (
	OK      = 00000 // 正常
	UNKNOWN = 00001 // 未知异常

	UserNotExit = 00001 // 用户不存在
)

var ServiceCode = map[int]string{
	OK:          "正常",
	UNKNOWN:     "未知异常",
	UserNotExit: "用户不存在",
}
