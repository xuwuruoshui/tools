package repository

var (
	OK = 00000			//
	UNKNOWN  = 00001 // 未知异常

	DateNoExit = 00001 // 不存在该数据
)

var RepoCode = map[int]string {
	OK:         "",
	UNKNOWN:    "未知异常",
	DateNoExit: "不存在该数据",
}
