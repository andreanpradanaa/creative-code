package controllers

import "github.com/labstack/echo"

type UserControllerInterfaces interface {
	Register(ctx echo.Context) error
}
