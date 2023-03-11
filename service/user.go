package service

import "github.com/FUnigrad/funiverse-workspace-service/model"

type IUserService interface {
	Get() []model.User
	Add(model.User) model.User
}

type UserService struct {
	User []model.User
}

func NewUserService() IUserService {
	return &UserService{}
}

func (service *UserService) Get() []model.User {
	return service.User
}

func (service *UserService) Add(user model.User) model.User {
	service.User = append(service.User, user)
	return user
}
