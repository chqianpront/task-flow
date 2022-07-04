package taskgroup

type TaskGroup struct {
	Name   string
	Id     int
	TaskId int
}

func NewTaskGroup() *TaskGroup {
	taskGroup := &TaskGroup{}
	return taskGroup
}
