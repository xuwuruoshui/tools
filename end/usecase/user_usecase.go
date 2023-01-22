package usecase

import (
	"context"
	"end/domain"
	"errors"
	"time"
)


type UserUseCase struct {
	userRepository domain.UserRepository[domain.User]
	contextTimeout time.Duration
}

func NewUserusecase(userepository domain.UserRepository[domain.User],tm time.Duration) UserUseCase{
	return UserUseCase{
		userRepository: userepository,
		contextTimeout: tm,
	}
}


func(u *UserUseCase) GetUser(c context.Context, id string) (*domain.User,error){
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.userRepository.GetUser(ctx,id)
}

func(u *UserUseCase) GetUserList(c context.Context,condition domain.PageDomain[domain.User]) (*domain.ListDomain[domain.User],error){
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.userRepository.GetUserList(ctx,condition)
}
func(u *UserUseCase) CreateUser(c context.Context, user domain.User) (*domain.RowAffect,error){
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.userRepository.CreateUser(ctx,user)
}

func(u *UserUseCase) UpdateUser(c context.Context,  user domain.User) (*domain.RowAffect,error){
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	if user.Id==""{
		return &domain.RowAffect{},errors.New("参数不完整")
	}
	return u.userRepository.UpdateUser(ctx,user)
}

func(u *UserUseCase) DeleteUser(c context.Context, id string) (*domain.RowAffect,error){
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.userRepository.DeleteUser(ctx,id)
}

