package handler

import "end/service"

type ApiRespData struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data any `json:"data"`
}

// api层的
func ApiResp(code int,data any)*ApiRespData {
	res := &ApiRespData{Data: data}
	res.Code = code
	res.Msg = ApiCode[code]
	return res
}


// servcie返回的
func ApiResp2(serviceRes *service.ServiceResData)*ApiRespData {
	res := &ApiRespData{Data: serviceRes.Data}
	switch serviceRes.Code {
	case OK:
		res.Code = OK
	case UNKNOWN:
		res.Code = UNKNOWN
	default:
		res.Code = BizError
	}
	return res
}


