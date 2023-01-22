package repository

import (
	"context"
	"end/bootstrap"
	"end/domain"
	"gorm.io/gorm"
	"strings"
)

type BaseDao interface {
	domain.User
	GetId()string
}


type Repository[T any] interface {
	GetById(context.Context, string) (*T,error)
	GetAllList(context.Context, domain.PageDomain[T]) (*domain.ListDomain[T],error)
	Create(context.Context, T) (*domain.RowAffect,error)
	Update(context.Context, T) (*domain.RowAffect,error)
	Delete(context.Context, string) (*domain.RowAffect,error)
}

type BaseRepository[T BaseDao] struct {
	App *bootstrap.App
}

func(u *BaseRepository[T]) GetById(c context.Context, id string) (*T,error){

	db := u.App.DBClient.(*bootstrap.MySqlClient).DB
	var t T
	tx := db.First(&t, id)
	if tx.Error!=nil && tx.Error!=gorm.ErrRecordNotFound{
		return nil,tx.Error
	}
	return &t,nil
}

func (u *BaseRepository[T]) GetAllList(c context.Context, p domain.PageDomain[T]) (*domain.ListDomain[T], error) {
	db := u.App.DBClient.(*bootstrap.MySqlClient).DB

	var tList domain.ListDomain[T]
	var t T
	var ts []*T
	var total int64

	condition := db.Model(t).Count(&total)
	if condition.Error!=nil  && condition.Error!=gorm.ErrRecordNotFound{
		return nil,condition.Error
	}

	condition.Find(&ts)

	tList.List = ts
	tList.Total = total

	return &tList,nil
}

func (u *BaseRepository[T]) Create(c context.Context, t T) (*domain.RowAffect, error) {
	db := u.App.DBClient.(*bootstrap.MySqlClient).DB
	tx := db.Create(&t)
	if tx.Error!=nil  && tx.Error!=gorm.ErrRecordNotFound{
		return nil,tx.Error
	}

	return &domain.RowAffect{Id: t.GetId(), Affect: tx.RowsAffected},nil
}

func (u *BaseRepository[T]) Update(c context.Context, t T) (*domain.RowAffect, error) {
	db := u.App.DBClient.(*bootstrap.MySqlClient).DB

	tx := db.Updates(&t)
	if tx.Error!=nil  && tx.Error!=gorm.ErrRecordNotFound{
		return nil,tx.Error
	}
	return &domain.RowAffect{Id: t.GetId(), Affect: tx.RowsAffected},nil
}

func (u *BaseRepository[T]) Delete(c context.Context, idsStr string) (*domain.RowAffect, error) {
	db := u.App.DBClient.(*bootstrap.MySqlClient).DB

	ids := strings.Split(idsStr, ",")
	var t T
	tx := db.Delete(&t, ids)
	if tx.Error!=nil  && tx.Error!=gorm.ErrRecordNotFound{
		return nil,tx.Error
	}
	return &domain.RowAffect{Id:idsStr, Affect: tx.RowsAffected},nil
}


