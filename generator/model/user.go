package model


type UserCenter struct {
	Name string `json:"name,omitempty" form:"name" gorm:"type:varchar(30);not null;index:idx_name;"`
	Age  *int32 `json:"age,omitempty" form:"age" gorm:"type:int(8);"`
	Age1  int32 `json:"age1,omitempty" form:"age1" gorm:"type:int(8);"`
	Email string `json:"email,omitempty" form:"email" gorm:"type:varchar(30);not null;index:idx_email;"`
	Phone string `json:"phone,omitempty" form:"phone" gorm:"type:varchar(30);not null;index:idx_phone;"`
}

