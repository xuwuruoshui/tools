package service

import (
	"context"
	"end/model"
	"end/repository"
	"time"
)


type UserService interface {
	GetUserById(context.Context, string) *ServiceResData
	GetUserAllList(context.Context, model.PageDomain[model.User]) *ServiceResData
	InsertUser(context.Context, model.User) *ServiceResData
	UpdateUser(context.Context, model.User) *ServiceResData
	DeleteUser(context.Context, string) *ServiceResData
	GetUserList(context.Context, model.PageDomain[model.User])*ServiceResData
}

type UserServiceImpl struct {
	userRepository repository.UserRepository[model.User]
	contextTimeout time.Duration
}

func NewUserService(userepository repository.UserRepository[model.User], tm time.Duration) UserServiceImpl {
	return UserServiceImpl{
		userRepository: userepository,
		contextTimeout: tm,
	}
}


func (u *UserServiceImpl) GetUserById(c context.Context, id string) *ServiceResData {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return ServiceResp2(u.userRepository.GetById(ctx, id))
}

func (u *UserServiceImpl) GetUserAllList(c context.Context, condition model.PageDomain[model.User]) *ServiceResData {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return ServiceResp2(u.userRepository.GetAllList(ctx, condition))
}

func (u *UserServiceImpl) InsertUser(c context.Context, user model.User) *ServiceResData {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return ServiceResp2(u.userRepository.Insert(ctx, user))
}

func (u *UserServiceImpl) UpdateUser(c context.Context, user model.User) *ServiceResData {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	return ServiceResp2(u.userRepository.Update(ctx, user))
}

func (u *UserServiceImpl) DeleteUser(c context.Context, id string) *ServiceResData {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return ServiceResp2(u.userRepository.Delete(ctx, id))
}

func (u *UserServiceImpl) GetUserList(c context.Context, condition model.PageDomain[model.User]) *ServiceResData {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return ServiceResp2(u.userRepository.GetList(ctx, condition))
}
