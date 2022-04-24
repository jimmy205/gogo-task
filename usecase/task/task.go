package taskUsecase

import (
	responseDto "gogolook/dto/response"
	taskDto "gogolook/dto/task"
	taskService "gogolook/service/task"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type TaskUsecase struct {
	taskService taskService.ITaskService
}

func NewTaskUsecase(
	taskService taskService.ITaskService,
) ITaskUsecase {
	return TaskUsecase{
		taskService: taskService,
	}
}

func (uc TaskUsecase) GetTasks(ctx *gin.Context) {

	tasks, err := uc.taskService.GetTasks()
	if err != nil {
		responseDto.FailWithError(ctx, err)
		return
	}

	taskDto := make([]taskDto.Task, 0, len(tasks))
	if err := copier.Copy(&taskDto, &tasks); err != nil {
		responseDto.FailWithError(ctx, err)
		return
	}

	responseDto.SuccessWithData(ctx, taskDto)
}

func (uc TaskUsecase) AddTask(ctx *gin.Context) {
	input := taskDto.AddTaskInput{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		responseDto.FailWithParams(ctx)
		return
	}

	newTask, err := uc.taskService.AddTask(input.Name)
	if err != nil {
		responseDto.FailWithError(ctx, err)
		return
	}

	taskDto := taskDto.Task{}
	if err := copier.Copy(&taskDto, &newTask); err != nil {
		responseDto.FailWithError(ctx, err)
		return
	}

	responseDto.SuccessWithData(ctx, taskDto)
}

func (uc TaskUsecase) EditTask(ctx *gin.Context) {
	input := taskDto.EditTaskInput{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		responseDto.FailWithParams(ctx)
		return
	}

	newTask, err := uc.taskService.EditTask(input.Id, input.Name, input.Status)
	if err != nil {
		responseDto.FailWithError(ctx, err)
		return
	}

	taskDto := taskDto.Task{}
	if err := copier.Copy(&taskDto, &newTask); err != nil {
		responseDto.FailWithError(ctx, err)
		return
	}

	responseDto.SuccessWithData(ctx, taskDto)
}

func (uc TaskUsecase) DeleteTask(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		responseDto.FailWithParams(ctx)
		return
	}

	if err := uc.taskService.DeleteTask(id); err != nil {
		responseDto.FailWithError(ctx, err)
		return
	}

	responseDto.Success(ctx)
}
