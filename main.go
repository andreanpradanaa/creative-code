package main

import (
	"creative-write/config"
	"creative-write/controllers"
	repository "creative-write/repositories"
	"creative-write/services"

	"github.com/labstack/echo"
)

func main() {

	DB := config.NewDB()
	userRepo := repository.NewUserRepositoryImpl(DB)
	userService := services.NewUserServiceImpl(userRepo)
	userController := controllers.NewUserControllerImpl(userService)

	e := echo.New()
	e.POST("/register", userController.Register)
	e.POST("/login", userController.Login)

	e.Logger.Fatal(e.Start(":8000"))
}
