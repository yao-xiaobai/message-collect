package main

import (
	"flag"
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
	o, err := gatherOptions(
		flag.NewFlagSet(os.Args[0], flag.ExitOnError),
		os.Args[1:]...,
	)
	if err != nil {
		logrus.Fatalf("new Options failed, err:%s", err.Error())
	}

	cfg := new(config.Config)
	logrus.Info(os.Args[1:])
	if err := utils.LoadFromYaml(o.Config, cfg); err != nil {
		logrus.Error("Config初始化失败, err:", err)
		return nil
	}
	config.InitEurBuildConfig(o.EurBuildConfig)
	return cfg
}

func gatherOptions(fs *flag.FlagSet, args ...string) (Options, error) {
	var o Options
	o.AddFlags(fs)
	err := fs.Parse(args)

	return o, err
}

type Options struct {
	Config         string
	EurBuildConfig string
}

func (o *Options) AddFlags(fs *flag.FlagSet) {
	fs.StringVar(&o.Config, "config-file", "", "Path to config file.")
	fs.StringVar(&o.EurBuildConfig, "eur-build-config-file", "", "Path to eur-build config file.")
}
