package config

import (
	"github.com/opensourceways/message-collect/common/kafka"
	"github.com/opensourceways/message-collect/utils"
	"github.com/sirupsen/logrus"
)

var EurBuildConfigInstance EurBuildConfig

type EurBuildConfig struct {
	Consume kafka.ConsumeConfig `yaml:"consume"`
	Publish string              `yaml:"publish"`
}

func InitEurBuildConfig(configFile string) {
	cfg := new(EurBuildConfig)
	if err := utils.LoadFromYaml(configFile, cfg); err != nil {
		logrus.Error("Config初始化失败, err:", err)
		return
	}
	EurBuildConfigInstance = *cfg
}
