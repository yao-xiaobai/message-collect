package main

import (
	"github.com/opensourceways/message-collect/common/kafka"
	"github.com/opensourceways/message-collect/config"
	"github.com/opensourceways/message-collect/manager"
	"github.com/opensourceways/message-collect/plugin"
	"github.com/opensourceways/message-collect/utils"
	"github.com/opensourceways/server-common-lib/logrusutil"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	logrusutil.ComponentInit("message-collect")
	log := logrus.NewEntry(logrus.StandardLogger())
	cfg := Init()
	logrus.Info("start init kafka,address=" + cfg.Kafka.Address)
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
	logrus.Info(os.Args[1:])
	if err := utils.LoadFromYaml("/vault/secrets/conf.yaml", cfg); err != nil {
		logrus.Error("Config初始化失败, err:", err)
		return nil
	}
	return cfg
}
