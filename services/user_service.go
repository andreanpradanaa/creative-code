package services

import "creative-write/model/web"

type UserServiceInterfaces interface {
	Register(request web.UserRegisterRequest) error
}
