package taskgroup

import (
	"github.com/gin-gonic/gin"
)

type TaskGroup struct {
	Name   string
	Id     int
	TaskId int
}

func newTaskGroup(ctx *gin.Context) {
}
func RegistRoute(r *gin.RouterGroup) {
	r.POST("/create", newTaskGroup)
}
