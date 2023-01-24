package repository

import (
	"context"
	"end/bootstrap"
	"end/model"
	"gorm.io/gorm"
)

type UserRepository[T model.User] interface {
	// 实现通用Mapper
	Repository[T]

	// 自己实现的接口
	GetList(context.Context, model.PageDomain[T])*RepoResData
}

type userRepository struct {
	BaseRepository[model.User]
}

func NewUserRepository(app *bootstrap.App) UserRepository[model.User] {
	repository := &userRepository{}
	repository.App = app
	return repository
}

func (r userRepository) GetList(ctx context.Context,p model.PageDomain[model.User])*RepoResData{
	db := r.App.DBClient.(*bootstrap.MySqlClient).DB

	var tList model.ListDomain[model.User]
	var ts []*model.User
	var total int64

	condition := db.Model(model.User{})


	if p.Condition.Id!=""{
		condition.Where("id = ?",p.Condition.Id)
	}
	if p.Condition.Age!=nil{
		condition.Where("age = ?",p.Condition.Age)
	}
	if p.Condition.Phone!=""{
		condition.Where("phone = ?",p.Condition.Phone)
	}
	if p.Condition.Email!=""{
		condition.Where("email = ?",p.Condition.Email)
	}
	condition.Count(&total)
	if condition.Error!=nil  && condition.Error!=gorm.ErrRecordNotFound{
		return RepoResp(UNKNOWN,condition.Error)
	}
	condition.Scopes(model.Paginate(p.PageNo,p.PageSize)).Find(&ts)

	tList.List = ts
	tList.Total = total

	return RepoResp(OK,&tList)
}