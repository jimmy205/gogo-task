package taskRepository

import taskModel "gogolook/model/task"

type ITaskRepository interface {
	GetTasks() ([]taskModel.Task, error)
	AddTask(name string) (taskModel.Task, error)
	EditTask(Id int, name string, status int) (taskModel.Task, error)
	DeleteTask(Id int) error
}
