package service

import (
	 "end/repository"
)

type ServiceResData struct {
	Code int `json:"code,omitempty"`
	Msg string `json:"msg,omitempty"`
	Data any `json:"data,omitempty"`
}

// service自身的
func ServiceResp(code int,data any)*ServiceResData {
	res := &ServiceResData{Data: data}
	res.Code = code
	res.Msg = ServiceCode[code]
	return res
}

// 持久层的
func ServiceResp2(repoRes *repository.RepoResData)*ServiceResData {
	res := &ServiceResData{Data: repoRes.Data}
	switch repoRes.Code {
	case OK:
		res.Code = OK
	case UNKNOWN:
		res.Code = UNKNOWN
	}
	return res
}

