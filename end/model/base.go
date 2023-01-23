package model

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
	PageNo int `json:"pageNo,omitempty"`
	PageSize int `json:"pageSize,omitempty"`
	Condition *T `json:"condition,omitempty"`
}

type ListDomain[T any] struct {
	Total int64 `json:"total,omitempty"`
	List []*T `json:"list,omitempty"`
}



type RowAffect[T string | []string] struct {
	Id T `json:"id,omitempty"`
	Affect int64 `json:"affect,omitempty"`
}

func Paginate(pageNo, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// 默认第一页
		if pageNo == 0 {
			pageNo = 1
		}

		// 最大页数100,默认10
		if pageSize > 100 {
			pageSize = 100
		} else if pageSize <= 0 {
			pageSize = 5
		}

		// 分页
		offset := (pageNo - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
