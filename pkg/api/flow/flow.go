package flow

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getFlow(ctx *gin.Context) {
	id := ctx.Param("id")
	log.Printf("get id: %s\n", id)
	ctx.String(http.StatusOK, "flow, %s", id)
}
func RegistRoute(r *gin.RouterGroup) {
	r.GET("/:id", getFlow)
}
