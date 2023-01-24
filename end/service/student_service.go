package service

import (
	"context"
	"end/model"
	"end/repository"
	"time"
)


type StudentService interface {
	GetStudentById(context.Context, string) *ServiceResData
	GetStudentAllList(context.Context, model.PageDomain[model.Student]) *ServiceResData
	InsertStudent(context.Context, model.Student) *ServiceResData
	UpdateStudent(context.Context, model.Student) *ServiceResData
	DeleteStudent(context.Context, string) *ServiceResData
	GetStudentList(context.Context, model.PageDomain[model.Student])*ServiceResData
}

type StudentServiceImpl struct {
	studentRepository repository.StudentRepository[model.Student]
	contextTimeout time.Duration
}

func NewStudentService(studentepository repository.StudentRepository[model.Student], tm time.Duration) StudentServiceImpl {
	return StudentServiceImpl{
		studentRepository: studentepository,
		contextTimeout: tm,
	}
}


func (u *StudentServiceImpl) GetStudentById(c context.Context, id string) *ServiceResData {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return ServiceResp2(u.studentRepository.GetById(ctx, id))
}

func (u *StudentServiceImpl) GetStudentAllList(c context.Context, condition model.PageDomain[model.Student]) *ServiceResData {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return ServiceResp2(u.studentRepository.GetAllList(ctx, condition))
}

func (u *StudentServiceImpl) InsertStudent(c context.Context, student model.Student) *ServiceResData {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return ServiceResp2(u.studentRepository.Insert(ctx, student))
}

func (u *StudentServiceImpl) UpdateStudent(c context.Context, student model.Student) *ServiceResData {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	return ServiceResp2(u.studentRepository.Update(ctx, student))
}

func (u *StudentServiceImpl) DeleteStudent(c context.Context, id string) *ServiceResData {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return ServiceResp2(u.studentRepository.Delete(ctx, id))
}

func (u *StudentServiceImpl) GetStudentList(c context.Context, condition model.PageDomain[model.Student]) *ServiceResData {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return ServiceResp2(u.studentRepository.GetList(ctx, condition))
}
