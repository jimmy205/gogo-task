package taskService

import (
	taskBo "gogolook/bo/task"
	taskRepository "gogolook/repository/task"

	"github.com/jinzhu/copier"
)

type TaskService struct {
	taskRepository taskRepository.ITaskRepository
}

func NewTaskService(
	taskRepository taskRepository.ITaskRepository,
) ITaskService {
	return TaskService{taskRepository: taskRepository}
}

func (s TaskService) GetTasks() ([]taskBo.Task, error) {
	tasks, err := s.taskRepository.GetTasks()
	if err != nil {
		return []taskBo.Task{}, err
	}

	tasksBo := make([]taskBo.Task, 0, len(tasks))
	if err := copier.Copy(&tasksBo, &tasks); err != nil {
		return tasksBo, err
	}
	return tasksBo, nil
}

func (s TaskService) AddTask(name string) (taskBo.Task, error) {
	newTask, err := s.taskRepository.AddTask(name)
	if err != nil {
		return taskBo.Task{}, err
	}
	taskBo := taskBo.Task{}
	if err := copier.Copy(&taskBo, &newTask); err != nil {
		return taskBo, err
	}
	return taskBo, nil
}

func (s TaskService) EditTask(id int, name string, status int) (taskBo.Task, error) {
	task, err := s.taskRepository.EditTask(id, name, status)
	if err != nil {
		return taskBo.Task{}, err
	}
	taskBo := taskBo.Task{}
	if err := copier.Copy(&taskBo, &task); err != nil {
		return taskBo, err
	}
	return taskBo, nil
}

func (s TaskService) DeleteTask(id int) error {
	if err := s.taskRepository.DeleteTask(id); err != nil {
		return err
	}
	return nil
}
