package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opensourceways/message-collect/common/kafka"
	"github.com/opensourceways/message-collect/manager"
	"github.com/opensourceways/message-collect/plugin"
)

func main() {
	kafka.Init()
	engine := gin.Default()

	manager.AddRoute(plugin.GiteeServerPlugin{Engine: engine})
	go func() {
		manager.StartConsume(plugin.EurBuildPlugin{})
	}()

	engine.Run(fmt.Sprintf(":%d", 8081))
	select {}
}
