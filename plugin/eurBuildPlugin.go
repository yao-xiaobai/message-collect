package plugin

import (
	"encoding/json"
	"time"

	"github.com/IBM/sarama"
	kfklib "github.com/opensourceways/kafka-lib/agent"
	"github.com/opensourceways/message-collect/common/kafka"
	"github.com/opensourceways/message-collect/config"
	"github.com/sirupsen/logrus"
)

type EurBuildPlugin struct {
}

func (p EurBuildPlugin) StartConsume() {
	h := EurGroupHandler{}
	kafka.ConsumeGroup(config.EurBuildConfigInstance.Consume, &h)
}

type EurGroupHandler struct{}

func (h EurGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h EurGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

const start_topic = "org.openEuler.prod.eur.build.start"
const end_topic = "org.openEuler.prod.eur.build.end"

/*
*

	接受到kafka消息，直接转发到另外一个kafka
*/
func (h EurGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		var raw EurBuildMessageRaw
		msgBodyErr := json.Unmarshal(message.Value, &raw)
		if msgBodyErr != nil {
			return msgBodyErr
		}
		logrus.Error(raw.Body.Pkg, raw.Body.Pkg == nil)
		if (raw.Topic == start_topic || raw.Topic == end_topic) && raw.Body.Pkg != nil {
			raw.Time = time.Now()
			err := kfklib.Publish(config.EurBuildConfigInstance.Publish, nil, message.Value)
			if err != nil {
				logrus.Error(err)
				return err
			}
			session.MarkMessage(message, "")
			logrus.Info("send eur build success")
		} else {
			logrus.Info("范围外的topic" + raw.Topic)

		}
	}
	return nil
}

type EurBuildMessageRaw struct {
	Body struct {
		User    string      `json:"user"`
		Copr    string      `json:"copr"`
		Owner   string      `json:"owner"`
		Pkg     interface{} `json:"pkg"`
		Build   int         `json:"build"`
		Chroot  string      `json:"chroot"`
		Version interface{} `json:"version"`
		Status  int         `json:"status"`
		IP      string      `json:"ip"`
		Who     string      `json:"who"`
		Pid     int         `json:"pid"`
		What    string      `json:"what"`
	} `json:"body"`
	Headers struct {
		OpenEulerMessagingSchema string `json:"openEuler_messaging_schema"`
		SentAt                   string `json:"sent-at"`
	} `json:"headers"`
	ID    string    `json:"id"`
	Topic string    `json:"topic"`
	Time  time.Time `json:"time"`
}
