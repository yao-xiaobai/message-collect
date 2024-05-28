package main

import (
	"fmt"
	"github.com/opensourceways/message-collect/common/kafka"
	"github.com/opensourceways/message-collect/config"
	"github.com/opensourceways/message-collect/manager"
	"github.com/opensourceways/message-collect/plugin"
	"github.com/opensourceways/message-collect/utils"
	"github.com/opensourceways/server-common-lib/logrusutil"
	"github.com/sirupsen/logrus"
)

func main() {
	logrusutil.ComponentInit("message-push")
	log := logrus.NewEntry(logrus.StandardLogger())
	cfg := Init()
	if err := kafka.Init(&cfg.Kafka, log, false); err != nil {
		logrus.Errorf("init kafka failed, err:%s", err.Error())
		return
	}
	go func() {
		manager.StartConsume(plugin.EurBuildPlugin{})
	}()
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
