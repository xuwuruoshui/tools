package repository


type RepoResData struct {
	Code int `json:"code,omitempty"`
	Msg string `json:"msg,omitempty"`
	Data any `json:"data,omitempty"`
}

func RepoResp(code int,data any)*RepoResData {
	res := &RepoResData{Data: data}
	res.Code = code
	res.Msg = RepoCode[code]
	return res
}

func (r RepoResData) SetMsg(msg string) RepoResData {
	r.Msg = msg
	return r
}