package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-collect/common/kafka"
	"message-collect/manager"
	"message-collect/plugin"
)

func main() {
	kafka.Init()
	engine := gin.Default()

	manager.AddRoute(plugin.GiteeServerPlugin{Engine: engine})
	manager.StartTask(plugin.EurBuildPlugin{})

	engine.Run(fmt.Sprintf(":%d", 8081))
	select {}
}
