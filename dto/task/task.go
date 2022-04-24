package taskDto

type Task struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

type AddTaskInput struct {
	Name string `json:"name" binding:"required"`
}

type EditTaskInput struct {
	Id     int    `json:"id" binding:"required"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

// type DeleteTaskInput struct {
// 	Id int `json:"id" binding:"required"`
// }
