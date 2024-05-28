package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opensourceways/message-collect/common/kafka"
	"github.com/opensourceways/message-collect/config"
	"github.com/opensourceways/message-collect/manager"
	"github.com/opensourceways/message-collect/plugin"
	"github.com/opensourceways/message-collect/utils"
	"github.com/opensourceways/server-common-lib/logrusutil"
)

func main() {
	logrusutil.ComponentInit("message-push")

	engine := gin.Default()
	cfg := Init()
	kafka.Init(*cfg)
	manager.AddRoute(plugin.GiteeServerPlugin{Engine: engine})
	go func() {
		manager.StartConsume(plugin.EurBuildPlugin{})
	}()
	engine.Run(fmt.Sprintf(":%d", cfg.Port))
	select {}
}

func Init() *config.Config {
	cfg := new(config.Config)
	if err := utils.LoadFromYaml("config/conf.yaml", cfg); err != nil {
		fmt.Println("Config初始化失败, err:", err)
		return nil
	}
	return cfg
}
