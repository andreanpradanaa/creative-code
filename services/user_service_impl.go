package services

import (
	"creative-write/config"
	"creative-write/model/domain"
	"creative-write/model/web"
	repository "creative-write/repositories"
	"creative-write/utils"
	"fmt"
)

type UserServiceImpl struct {
	UserRepo repository.UserRepositoryInterfaces
}

func NewUserServiceImpl(UserRepo repository.UserRepositoryInterfaces) UserServiceInterfaces {
	return &UserServiceImpl{
		UserRepo: UserRepo,
	}
}

func (service *UserServiceImpl) Login(request web.UserLoginRequest) (string, error) {
	user, err := service.UserRepo.Find(request.Email)
	if err != nil {
		return "", fmt.Errorf("user not Found")
	}

	err = utils.VerifyPassword(user.Password, request.Password)
	if err != nil {
		return "", err
	}

	cfg, _ := config.LoadConfig(".")
	fmt.Println(cfg)
	token, err := utils.GenerateToken(cfg.AccessTokenExpiresIn, user.Id, cfg.AccessTokenPrivateKey)
	fmt.Println("token", token)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (service *UserServiceImpl) Register(request web.UserRegisterRequest) error {

	user := domain.User{}
	user.Name = request.Name
	user.Email = request.Email
	hashPassword, err := utils.HashPassword(request.Password)
	fmt.Println(hashPassword)
	if err != nil {
		return err
	}
	user.Password = hashPassword

	if err := service.UserRepo.Save(&user); err != nil {
		return err
	}

	return nil
}
