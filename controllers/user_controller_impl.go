package controllers

import (
	"creative-write/model/web"
	"creative-write/services"
	"net/http"

	"github.com/labstack/echo"
)

type UserControllerImpl struct {
	UserService services.UserServiceInterfaces
}

func NewUserControllerImpl(UserService services.UserServiceInterfaces) UserControllerInterfaces {
	return &UserControllerImpl{
		UserService: UserService,
	}
}

func (controller *UserControllerImpl) Register(ctx echo.Context) error {
	userRegisterRequest := web.UserRegisterRequest{}

	if err := ctx.Bind(&userRegisterRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, "error register")
	}

	if err := controller.UserService.Register(userRegisterRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, "error register")
		return err
	}

	ctx.JSON(http.StatusOK, "register sukses")
	return nil
}
