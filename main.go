package main

import (
	"fmt"
	"log"
	"os"

	"chen.com/task-flow/pkg/api"
	"chen.com/task-flow/pkg/model"
	"chen.com/task-flow/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	osArgsLen := len(os.Args)
	if osArgsLen < 2 {
		log.Printf("exec args is not enough\n")
		return
	}
	actionType := os.Args[1]
	envrionmentType := os.Args[2]
	log.Printf("%s server environment %s,", actionType, envrionmentType)
	conf := utils.GetConf(envrionmentType)
	addrStr := fmt.Sprintf("%s:%d", conf.Server.Ip, conf.Server.Port)
	model.Init(utils.ConnDB(conf))
	route := gin.Default()
	v1 := route.Group("api/v1")
	{
		api.RegistRoute(v1)
	}
	route.Run(addrStr)
}
