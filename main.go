package main

import (
	"encoding/json"
	"flag"
	kfklib "github.com/opensourceways/kafka-lib/agent"
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
		Id        int    `json:"id"`
		Topic     string `json:"topic"`
		Community string `json:"community"`
		GroupName string `json:"group_name"`
		Sponsor   string `json:"sponsor"`
		Date      string `json:"date"`
		Start     string `json:"start"`
		End       string `json:"end"`
		Duration  string `json:"duration"`
		Agenda    string `json:"agenda"`
		Etherpad  string `json:"etherpad"`
		Emaillist string `json:"emaillist"`
		HostId    string `json:"host_id"`
		Mid       string `json:"mid"`
		Mmid      string `json:"mmid"`
		JoinUrl   string `json:"join_url"`
		IsDelete  int    `json:"is_delete"`
		StartUrl  string `json:"start_url"`
		Timezone  string `json:"timezone"`
		User      int    `json:"user"`
		Group     int    `json:"group"`
		Mplatform string `json:"mplatform"`
	} `json:"msg"`
}

func PublishMeeting() {
	message := `
	{
  "action": "create_meeting",
  "msg": {
    "id": 2413,
    "topic": "哈哈哈还",
    "community": "openeuler",
    "group_name": "Infrastructure",
    "sponsor": "shishupei10",
    "date": "2024-08-24",
    "start": "08:00",
    "end": "08:30",
    "duration": null,
    "agenda": "还吐饿一天没人理我",
    "etherpad": "https://etherpad.openeuler.org/p/p/Infrastructure-meetings",
    "emaillist": "infra@openeuler.org;",
    "host_id": "iGNZaz-9SiCleuoyYesRJg",
    "mid": "89199048779",
    "mmid": null,
    "join_url": "https://us06web.zoom.us/j/89199048779?pwd\u003djUdhKq9WCJEvai2LqvVtsmDj5LIjk0.1",
    "is_delete": 0,
    "start_url": "https://us06web.zoom.us/s/89199048779?zak\u003deyJ0eXAiOiJKV1QiLCJzdiI6IjAwMDAwMSIsInptX3NrbSI6InptX28ybSIsImFsZyI6IkhTMjU2In0.eyJpc3MiOiJ3ZWIiLCJjbHQiOjAsIm1udW0iOiI4OTE5OTA0ODc3OSIsImF1ZCI6ImNsaWVudHNtIiwidWlkIjoiaUdOWmF6LTlTaUNsZXVveVllc1JKZyIsInppZCI6IjBmYmMzZTQ0MjY5NTQwODQ4MmQwYjFiMGU4ZjUyY2MzIiwic2siOiIwIiwic3R5IjoxLCJ3Y2QiOiJ1czA2IiwiZXhwIjoxNzI0MTQ4NTkwLCJpYXQiOjE3MjQxNDEzOTAsImFpZCI6IktTUDFGekRNUjh5elNmSzRyRHhOYlEiLCJjaWQiOiIifQ.EKV8oujz5UM9qF3zBznS1kP9-uilpp2hUfHfiw_tYog",
    "timezone": "Asia/Shanghai",
    "user": 515,
    "group": 10,
    "mplatform": "zoom"
  }
}
`
	var raw OpenEulerMeetingRaw
	json.Unmarshal([]byte(message), &raw)
	value, _ := json.Marshal(raw)
	err := kfklib.Publish("openEuler_meeting_raw", nil, value)
	if err != nil {
		return
	}
}

func main() {
	logrusutil.ComponentInit("message-collect")
	log := logrus.NewEntry(logrus.StandardLogger())
	cfg := Init()
	if err := kafka.Init(&cfg.Kafka, log, false); err != nil {
		logrus.Errorf("init kafka failed, err:%s", err.Error())
		return
	}
	//PublishMeeting()

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
