package usecase

import "app/model"

type IUserUseCase interface {
	SignUp(user *model.User) (model.UserResponse, error)
	Login(user *model.User) (string, error)
}
