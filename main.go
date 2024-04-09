package main

import (
	"flag"
	"fmt"
	"github.com/opensourceways/server-common-lib/logrusutil"
	liboptions "github.com/opensourceways/server-common-lib/options"
	"github.com/sirupsen/logrus"
	"message-collect/common/kafka"
	"message-collect/config"
	"message-collect/models/domain/event"
	"message-collect/models/messageadapter"
	"message-collect/service/collector"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	logrusutil.ComponentInit("messageAdapter-collect")
	log := logrus.NewEntry(logrus.StandardLogger())

	cfg := new(config.Config)
	// kafka
	if err := kafka.Init(&cfg.Kafka, log, false); err != nil {
		logrus.Errorf("init kafka failed, err:%s", err.Error())
		return
	}
	defer kafka.Exit()
	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(time.Second * 5)
			publish()
			fmt.Println("发送一条新消息")
		}
	}()

	go func() {
		collector.Consume()
	}()
	<-sig

}

func publish() {
	e := event.NewEurBuildEvent()
	if err1 := messageadapter.SendMsg("eur_build_raw", &e); err1 != nil {
		logrus.Errorf("")
	}
}
