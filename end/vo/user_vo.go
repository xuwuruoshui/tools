package vo

type UserReq struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Age  *int32 `json:"age,omitempty"`
}

type UserResp struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Age  *int32 `json:"age,omitempty"`
}
