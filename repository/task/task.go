package taskRepository

import (
	"fmt"
	taskModel "gogolook/model/task"
	"sort"
	"sync"
)

type TaskRepository struct {
	id      *idCreator
	dataMap map[int]taskModel.Task
	mLock   *sync.RWMutex
}

func NewTaskRepository() TaskRepository {
	return TaskRepository{
		id:      &idCreator{},
		dataMap: make(map[int]taskModel.Task),
		mLock:   new(sync.RWMutex),
	}
}

type idCreator struct {
	id int
}

func (c *idCreator) add() {
	c.id++
}

func (c idCreator) get() int {
	return c.id
}

func (r TaskRepository) GetTasks() ([]taskModel.Task, error) {
	r.mLock.RLock()
	defer r.mLock.RUnlock()

	tasks := []taskModel.Task{}
	for _, task := range r.dataMap {
		tasks = append(tasks, task)
	}

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Id > tasks[j].Id
	})

	return tasks, nil
}

func (r TaskRepository) AddTask(name string) (taskModel.Task, error) {
	r.mLock.Lock()
	defer r.mLock.Unlock()

	r.id.add()

	r.dataMap[r.id.get()] = taskModel.Task{
		Id:     r.id.get(),
		Name:   name,
		Status: 0, // default incomplete
	}

	return r.dataMap[r.id.get()], nil
}

func (r TaskRepository) EditTask(id int, name string, status int) (taskModel.Task, error) {
	r.mLock.Lock()
	defer r.mLock.Unlock()

	task, ok := r.dataMap[id]
	if !ok {
		return taskModel.Task{}, fmt.Errorf("id %d not found", id)
	}
	if name != "" {
		task.Name = name
	}
	if status != 0 {
		task.Status = status
	}

	r.dataMap[id] = task
	return r.dataMap[id], nil
}

func (r TaskRepository) DeleteTask(id int) error {
	r.mLock.Lock()
	defer r.mLock.Unlock()

	_, ok := r.dataMap[id]
	if !ok {
		return fmt.Errorf("id %d not found", id)
	}

	delete(r.dataMap, id)
	return nil
}
