package config

import (
	"github.com/opensourceways/message-collect/common/kafka"
	"github.com/opensourceways/message-collect/utils"
	"github.com/sirupsen/logrus"
)

var OpenEulerMeetingConfigInstance OpenEulerMeetingConfig

type OpenEulerMeetingConfig struct {
	Consume kafka.ConsumeConfig `yaml:"consume"`
	Publish string              `yaml:"publish"`
}

func InitOpenEulerMeetingConfig(configFile string) {
	cfg := new(OpenEulerMeetingConfig)
	if err := utils.LoadFromYaml(configFile, cfg); err != nil {
		logrus.Error("Config初始化失败, err:", err)
		return
	}
	OpenEulerMeetingConfigInstance = *cfg
}
