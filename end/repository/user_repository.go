package repository

import (
	"context"
	"end/domain"
	"end/tools"
)

type UserRepository interface {
	GetUser(context.Context) domain.User
}

func (u *User) GetUser(ctx context.Context) domain.User {
	return User{
		Id:   "1",
		Name: "test",
		Age:  tools.ToPointer(int32(10)),
	}
}
