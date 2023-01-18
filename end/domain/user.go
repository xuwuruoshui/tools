package domain

import (
	"context"
)

type User struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Age  *int32 `json:"age,omitempty"`
}

type UserService interface {
	GetUser(ctx context.Context) User
}

func (u User) GetUser(ctx context.Context) User {

	// 逻辑

	// 调用repository获取数据


	return u
}
