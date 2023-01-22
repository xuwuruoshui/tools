package repository

import (
	"context"
	"end/bootstrap"
	"end/domain"
	"end/utils"
	"gorm.io/gorm"
	"strings"
)

type userRepository struct {
	App *bootstrap.App
}

func NewUserRepository(app *bootstrap.App)domain.UserRepository[domain.User] {
	return &userRepository{App: app}
}


func(u *userRepository) GetUser(c context.Context, id string) (*domain.User,error){

	db := u.App.DBClient.(*bootstrap.MySqlClient).DB
	user := &domain.User{}
	tx := db.First(user, id)
	if tx.Error!=nil && tx.Error!=gorm.ErrRecordNotFound{
		return nil,tx.Error
	}
	return user,nil
}

func (u *userRepository) GetUserList(c context.Context, p domain.PageDomain[domain.User]) (*domain.ListDomain[domain.User], error) {
	db := u.App.DBClient.(*bootstrap.MySqlClient).DB

	var userList domain.ListDomain[domain.User]
	var users []*domain.User
	var total int64

	tx := db.Model(domain.User{}).Count(&total).Offset(int(p.PageNo)).Limit(int(p.PageSize)).Find(&users)
	if tx.Error!=nil  && tx.Error!=gorm.ErrRecordNotFound{
		return nil,tx.Error
	}
	userList.List = users
	userList.Total = total

	return &userList,nil
}

func (u *userRepository) CreateUser(c context.Context, user domain.User) (*domain.RowAffect, error) {
	db := u.App.DBClient.(*bootstrap.MySqlClient).DB
	user.Id = utils.GenerateId()
	tx := db.Create(&user)
	if tx.Error!=nil  && tx.Error!=gorm.ErrRecordNotFound{
		return nil,tx.Error
	}
	return &domain.RowAffect{Id: user.Id, Affect: tx.RowsAffected},nil
}

func (u *userRepository) UpdateUser(c context.Context, user domain.User) (*domain.RowAffect, error) {
	db := u.App.DBClient.(*bootstrap.MySqlClient).DB

	tx := db.Updates(user)
	if tx.Error!=nil  && tx.Error!=gorm.ErrRecordNotFound{
		return nil,tx.Error
	}
	return &domain.RowAffect{Id: user.Id, Affect: tx.RowsAffected},nil
}

func (u *userRepository) DeleteUser(c context.Context, idsStr string) (*domain.RowAffect, error) {
	db := u.App.DBClient.(*bootstrap.MySqlClient).DB

	ids := strings.Split(idsStr, ",")
	tx := db.Delete(&domain.User{}, ids)
	if tx.Error!=nil  && tx.Error!=gorm.ErrRecordNotFound{
		return nil,tx.Error
	}
	return &domain.RowAffect{Id:idsStr, Affect: tx.RowsAffected},nil
}


