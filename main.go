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

type OpenEulerMeetingRaw struct {
	Action string `json:"action"`
	Msg    struct {
		Topic     string      `json:"topic"`
		Platform  interface{} `json:"platform"`
		Sponsor   string      `json:"sponsor"`
		GroupName string      `json:"group_name"`
		GroupId   interface{} `json:"group_id"`
		Date      string      `json:"date"`
		Start     string      `json:"start"`
		End       string      `json:"end"`
		Etherpad  string      `json:"etherpad"`
		Agenda    string      `json:"agenda"`
		EmailList string      `json:"emaillist"`
		Record    string      `json:"record"`
		JoinUrl   string      `json:"join_url"`
	} `json:"msg"`
}

func main() {
	logrusutil.ComponentInit("message-collect")
	log := logrus.NewEntry(logrus.StandardLogger())
	cfg := Init()
	if err := kafka.Init(&cfg.Kafka, log, false); err != nil {
		logrus.Errorf("init kafka failed, err:%s", err.Error())
		return
	}
	go func() {
		manager.StartConsume(plugin.EurBuildPlugin{})
	}()

	go func() {
		manager.StartConsume(plugin.OpenEulerMeetingPlugin{})
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
	config.InitOpenEulerMeetingConfig(o.MeetingConfig)
	return cfg
}

/*
获取启动参数，配置文件地址由启动参数传入
*/
func gatherOptions(fs *flag.FlagSet, args ...string) (Options, error) {
	var o Options
	o.AddFlags(fs)
	err := fs.Parse(args)

	return o, err
}

type Options struct {
	Config         string
	EurBuildConfig string
	MeetingConfig  string
}

func (o *Options) AddFlags(fs *flag.FlagSet) {
	fs.StringVar(&o.Config, "config-file", "", "Path to config file.")
	fs.StringVar(&o.EurBuildConfig, "eur-build-config-file", "", "Path to eur-build config file.")
	fs.StringVar(&o.MeetingConfig, "meeting-config-file", "", "Path to meeting config file.")

}
