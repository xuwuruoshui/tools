package domain

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	Id string `json:"id,omitempty" gorm:"primary_key"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	DeletedAt *gorm.DeletedAt `json:"deletedAt,omitempty" gorm:"index"`
}


type PageDomain[T any] struct {
	PageNo int32 `json:"pageNo,omitempty"`
	PageSize int32 `json:"pageSize,omitempty"`
	Condition T `json:"condition,omitempty"`
}

type ListDomain[T any] struct {
	Total int64 `json:"total,omitempty"`
	List []*T `json:"list,omitempty"`
}

type RowAffect struct {
	Id string `json:"id,omitempty"`
	Affect int64 `json:"affect,omitempty"`
}