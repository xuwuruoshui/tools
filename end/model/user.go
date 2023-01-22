package model

import (
	"end/utils"
	"gorm.io/gorm"
)

type User struct {
	BaseModel
	Name string `json:"name,omitempty" form:"name" gorm:"type:varchar(30);not null;index:idx_name;"`
	Age  *int32 `json:"age,omitempty" form:"age" gorm:"type:int(8);"`
	Email string `json:"email,omitempty" form:"email" gorm:"type:varchar(30);not null;index:idx_email;"`
	Phone string `json:"phone,omitempty" form:"phone" gorm:"type:varchar(30);not null;index:idx_phone;"`
}

func (u User) TableName() string{
	return "user"
}

func(u User) GetId()string{
	return u.Id
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = utils.GenerateId()
	return
}

