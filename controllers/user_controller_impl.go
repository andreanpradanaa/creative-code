package controllers

import (
	"creative-write/model/web"
	"creative-write/services"
	"fmt"
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

func (controller *UserControllerImpl) Login(ctx echo.Context) error {

	fmt.Println("controller")
	userLoginRequest := web.UserLoginRequest{}
	if err := ctx.Bind(&userLoginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, "invalid email or password")
	}

	token, err := controller.UserService.Login(userLoginRequest)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "wrong password")
		return err
	}

	// response := []string{
	// 	"sukses",
	// 	"token: ", token,
	// }
	ctx.SetCookie(
		&http.Cookie{
			Name:  "access_token",
			Value: token,
			// MaxAge:   config.AccessTokenMaxAge * 60, // Convert to seconds
			Path:     "/",
			Domain:   "localhost",
			Secure:   false,
			HttpOnly: true,
		},
	)
	ctx.JSON(http.StatusOK, token)
	return nil
}
