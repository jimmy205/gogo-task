package taskUsecase

import "github.com/gin-gonic/gin"

type ITaskUsecase interface {
	GetTasks(ctx *gin.Context)
	AddTask(ctx *gin.Context)
	EditTask(ctx *gin.Context)
	DeleteTask(ctx *gin.Context)
}
