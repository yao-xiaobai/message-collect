package main

import (
	"flag"
	"github.com/opensourceways/server-common-lib/logrusutil"
	liboptions "github.com/opensourceways/server-common-lib/options"
	"github.com/sirupsen/logrus"
	"message-collect/common/kafka"
	"message-collect/config"
	"message-collect/models/domain/event"
	"message-collect/models/messageadapter"
	"message-collect/service/collector"
)

type options struct {
	service     liboptions.ServiceOptions
	enableDebug bool
}

func (o *options) Validate() error {
	return o.service.Validate()
}

func gatherOptions(fs *flag.FlagSet, args ...string) (options, error) {
	var o options

	o.service.AddFlags(fs)

	fs.BoolVar(
		&o.enableDebug, "enable_debug", false,
		"whether to enable debug model.",
	)

	err := fs.Parse(args)

	return o, err
}

func main() {
	logrusutil.ComponentInit("messageAdapter-collect")
	log := logrus.NewEntry(logrus.StandardLogger())

	cfg := new(config.Config)
	// kafka
	if err := kafka.Init(&cfg.Kafka, log, false); err != nil {
		logrus.Errorf("init kafka failed, err:%s", err.Error())
		return
	}
	defer kafka.Exit()
	publish()
	//publish()
	collector.Consume()
}

func publish() {
	e := event.NewEurBuildEvent()
	if err1 := messageadapter.SendMsg(&e); err1 != nil {
		logrus.Errorf("")
	}
}
