package taskService

import taskBo "gogolook/bo/task"

type ITaskService interface {
	GetTasks() ([]taskBo.Task, error)
	AddTask(name string) (taskBo.Task, error)
	EditTask(Id int, name string, status int) (taskBo.Task, error)
	DeleteTask(Id int) error
}
