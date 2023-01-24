package repository

import (
	"context"
	"end/bootstrap"
	"end/model"
	"gorm.io/gorm"
)

type StudentRepository[T model.Student] interface {
	// 实现通用Mapper
	Repository[T]

	// 自己实现的接口
	GetList(context.Context, model.PageDomain[T])*RepoResData
}

type studentRepository struct {
	BaseRepository[model.Student]
}

func NewStudentRepository(app *bootstrap.App) StudentRepository[model.Student] {
	repository := &studentRepository{}
	repository.App = app
	return repository
}

func (r studentRepository) GetList(ctx context.Context,p model.PageDomain[model.Student])*RepoResData{
	db := r.App.DBClient.(*bootstrap.MySqlClient).DB

	var tList model.ListDomain[model.Student]
	var ts []*model.Student
	var total int64

	condition := db.Model(model.Student{})

    if p.Condition.Username!=""{
    		condition.Where("_username = ?",p.Condition.Username)
    }
    if p.Condition.Age!=nil{
    		condition.Where("_age = ?",p.Condition.Age)
    }
    if p.Condition.Email!=""{
    		condition.Where("_email = ?",p.Condition.Email)
    }
    if p.Condition.Phone!=""{
    		condition.Where("_phone = ?",p.Condition.Phone)
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