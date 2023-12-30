package services

import "creative-write/model/web"

type UserServiceInterfaces interface {
	Register(request web.UserRegisterRequest) error
	Login(request web.UserLoginRequest) (string, error)
}
