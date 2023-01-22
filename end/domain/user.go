package domain

import (
	"context"
)

type User struct {
	BaseModel
	Name string `json:"name,omitempty" gorm:"type:varchar(30);not null;index:idx_name;"`
	Age  *int32 `json:"age,omitempty" gorm:"type:int(8);"`
	Email string `json:"email,omitempty" gorm:"type:varchar(30);not null;index:idx_email;"`
	Phone string `json:"phone,omitempty" gorm:"type:varchar(30);not null;index:idx_phone;"`
}

type UserRepository[T User] interface {
	GetUser(context.Context, string) (*User,error)
	GetUserList(context.Context, PageDomain[T]) (*ListDomain[T],error)
	CreateUser(context.Context, User) (*RowAffect,error)
	UpdateUser(context.Context, User) (*RowAffect,error)
	DeleteUser(context.Context, string) (*RowAffect,error)
}

