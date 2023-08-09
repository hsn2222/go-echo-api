package main

import (
	"react_go_echo/controller"
	"react_go_echo/db"
	"react_go_echo/repository"
	"react_go_echo/validator"
	"react_go_echo/usecase"
	"react_go_echo/router"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)

	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()

	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)

	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)

	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))
}