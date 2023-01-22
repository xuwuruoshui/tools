package repository

import (
	"context"
	"end/bootstrap"
	"end/domain"
	"gorm.io/gorm"
)

type UserRepository[T domain.User] interface {
	// 实现通用Mapper
	Repository[T]

	// 自己实现的接口
	GetList(context.Context,domain.PageDomain[T])(*domain.ListDomain[T],error)
}

type userRepository struct {
	BaseRepository[domain.User]
}

func NewUserRepository(app *bootstrap.App) UserRepository[domain.User] {
	repository := &userRepository{}
	repository.App = app
	return repository
}

func (u userRepository) GetList(ctx context.Context,p domain.PageDomain[domain.User])(*domain.ListDomain[domain.User],error){
	db := u.App.DBClient.(*bootstrap.MySqlClient).DB

	var tList domain.ListDomain[domain.User]
	var ts []*domain.User
	var total int64

	condition := db.Model(domain.User{})


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
		return nil,condition.Error
	}
	condition.Scopes(domain.Paginate(p.PageNo,p.PageSize)).Find(&ts)

	tList.List = ts
	tList.Total = total

	return &tList,nil
}