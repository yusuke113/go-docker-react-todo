package main

import (
	"app/controller"
	"app/db"
	"app/repository"
	"app/router"
	"app/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepository)
	taskUseCase := usecase.NewTaskUseCase(taskRepository)
	userController := controller.NewUserController(userUseCase)
	TaskController := controller.NewTaskController(taskUseCase)
	e := router.NewRouter(userController, TaskController)
	e.Logger.Fatal(e.Start(":8080"))
}
