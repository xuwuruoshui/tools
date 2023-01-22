package usecase

import (
	"context"
	"end/domain"
	"end/repository"
	"errors"
	"time"
)


type UserUseCase struct {
	userRepository repository.UserRepository[domain.User]
	contextTimeout time.Duration
}

func NewUserusecase(userepository repository.UserRepository[domain.User], tm time.Duration) UserUseCase {
	return UserUseCase{
		userRepository: userepository,
		contextTimeout: tm,
	}
}

func(u *UserUseCase) GetById(c context.Context, id string) (*domain.User,error){
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.userRepository.GetById(ctx,id)
}

func(u *UserUseCase) GetAllList(c context.Context,condition domain.PageDomain[domain.User]) (*domain.ListDomain[domain.User],error){
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.userRepository.GetAllList(ctx,condition)
}

func(u *UserUseCase) Create(c context.Context, user domain.User) (*domain.RowAffect,error){
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.userRepository.Create(ctx,user)
}

func(u *UserUseCase) Update(c context.Context,  user domain.User) (*domain.RowAffect,error){
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	if user.Id==""{
		return &domain.RowAffect{},errors.New("参数不完整")
	}
	return u.userRepository.Update(ctx,user)
}

func(u *UserUseCase) Delete(c context.Context, id string) (*domain.RowAffect,error){
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.userRepository.Delete(ctx,id)
}

func(u *UserUseCase) GetList(c context.Context,condition domain.PageDomain[domain.User]) (*domain.ListDomain[domain.User],error){
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.userRepository.GetList(ctx, condition)
}
