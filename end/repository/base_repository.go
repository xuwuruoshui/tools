package repository

import (
	"context"
	"end/bootstrap"
	"end/model"
	"gorm.io/gorm"
	"strings"
)

type BaseDao interface {
	model.User
	GetId()string
}


type Repository[T any] interface {
	GetById(context.Context, string) *RepoResData
	GetAllList(context.Context, model.PageDomain[T]) *RepoResData
	Create(context.Context, T) *RepoResData
	Update(context.Context, T) *RepoResData
	Delete(context.Context, string) *RepoResData
}

type BaseRepository[T BaseDao] struct {
	App *bootstrap.App
}

func(u *BaseRepository[T]) GetById(c context.Context, id string) *RepoResData{

	db := u.App.DBClient.(*bootstrap.MySqlClient).DB
	var t T
	tx := db.First(&t, id)
	if tx.Error!=nil && tx.Error!=gorm.ErrRecordNotFound{

		return RepoResp(UNKNOWN,tx.Error)
	}
	return RepoResp(OK,t)
}

func (u *BaseRepository[T]) GetAllList(c context.Context, p model.PageDomain[T]) *RepoResData{
	db := u.App.DBClient.(*bootstrap.MySqlClient).DB

	var tList model.ListDomain[T]
	var t T
	var ts []*T
	var total int64

	tx := db.Model(t).Count(&total)
	if tx.Error!=nil  && tx.Error!=gorm.ErrRecordNotFound{
		return RepoResp(UNKNOWN,tx.Error)
	}

	tx.Find(&ts)

	tList.List = ts
	tList.Total = total

	return RepoResp(OK,&tList)
}

func (u *BaseRepository[T]) Create(c context.Context, t T)*RepoResData {
	db := u.App.DBClient.(*bootstrap.MySqlClient).DB
	tx := db.Create(&t)
	if tx.Error!=nil  && tx.Error!=gorm.ErrRecordNotFound{
		return RepoResp(UNKNOWN,tx.Error)
	}

	return RepoResp(OK, &model.RowAffect{Id: t.GetId(), Affect: tx.RowsAffected})
}

func (u *BaseRepository[T]) Update(c context.Context, t T) *RepoResData {
	db := u.App.DBClient.(*bootstrap.MySqlClient).DB

	tx := db.Updates(&t)
	if tx.Error!=nil  && tx.Error!=gorm.ErrRecordNotFound{
		return RepoResp(UNKNOWN,tx.Error)
	}
	return RepoResp(OK,&model.RowAffect{Id: t.GetId(), Affect: tx.RowsAffected})
}

func (u *BaseRepository[T]) Delete(c context.Context, idsStr string) *RepoResData {
	db := u.App.DBClient.(*bootstrap.MySqlClient).DB

	ids := strings.Split(idsStr, ",")
	var t T
	tx := db.Delete(&t, ids)
	if tx.Error!=nil  && tx.Error!=gorm.ErrRecordNotFound{
		return RepoResp(UNKNOWN,tx.Error)
	}
	return RepoResp(OK,&model.RowAffect{Id: idsStr, Affect: tx.RowsAffected})
}


