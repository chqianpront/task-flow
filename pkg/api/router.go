package api

import (
	"github.com/gin-gonic/gin"
)

type FlowRouter struct{}

func RegistRoute(r *gin.RouterGroup) {
	user := r.Group("user")
	login := r.Group("login")
	user.GET("list", ListUser)
	login.GET("login", Login)
}
