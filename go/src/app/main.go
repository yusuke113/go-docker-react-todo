package main

import (
	"app/controller"
	"app/db"
	"app/repository"
	"app/router"
	"app/usecase"
	"app/validator"
)

func main() {
	db := db.NewDB()
	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepository, userValidator)
	taskUseCase := usecase.NewTaskUseCase(taskRepository, taskValidator)
	userController := controller.NewUserController(userUseCase)
	TaskController := controller.NewTaskController(taskUseCase)
	e := router.NewRouter(userController, TaskController)
	e.Logger.Fatal(e.Start(":8080"))
}
