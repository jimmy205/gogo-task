package main

import (
	taskRepository "gogolook/repository/task"
	taskService "gogolook/service/task"
	taskUsecase "gogolook/usecase/task"

	"github.com/gin-gonic/gin"
)

func main() {

	taskRepository := taskRepository.NewTaskRepository()
	taskService := taskService.NewTaskService(taskRepository)
	taskUsecase := taskUsecase.NewTaskUsecase(taskService)

	r := gin.Default()

	startTask(r, taskUsecase)

	r.Run(":8000")
}

func startTask(
	r *gin.Engine,
	usercase taskUsecase.ITaskUsecase,
) {

	r.GET("/tasks", usercase.GetTasks)
	r.POST("/task", usercase.AddTask)
	r.PUT("/task/:id", usercase.EditTask)
	r.DELETE("/task/:id", usercase.DeleteTask)
}
