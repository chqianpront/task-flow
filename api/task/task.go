package task

import "chen.com/task-flow/api/taskgroup"

type task struct {
	Name      string
	CreatorId int
}
type MTaskGroup struct {
	*taskgroup.TaskGroup
}

func NewTask() *task {
	t := &task{
		Name: "haha",
	}
	t.addCreator(222)
	return t
}
func (mtg *MTaskGroup) AddName(name string) {
	mtg.Name = name
}
func NewMTaskGroup() *MTaskGroup {
	mtg := &MTaskGroup{}
	mtg.AddName("test")
	return mtg
}
