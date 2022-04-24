package taskModel

type Task struct {
	Id     int    `gorm:"column:id"`
	Name   string `gorm:"column:name"`
	Status int    `gorm:"column:status"`
}
