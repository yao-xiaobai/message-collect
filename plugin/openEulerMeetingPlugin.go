package plugin

import (
	"encoding/json"
	"github.com/IBM/sarama"
	kfklib "github.com/opensourceways/kafka-lib/agent"
	"github.com/opensourceways/message-collect/common/kafka"
	"github.com/opensourceways/message-collect/config"
	"github.com/sirupsen/logrus"
)

type OpenEulerMeetingPlugin struct {
}

func (p OpenEulerMeetingPlugin) StartConsume() {
	h := OpenEulerMeetingHandler{}
	kafka.ConsumeGroup(config.OpenEulerMeetingConfigInstance.Consume, &h)
}

type OpenEulerMeetingHandler struct{}

func (h OpenEulerMeetingHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h OpenEulerMeetingHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

/*
*

	接受到kafka消息，直接转发到另外一个kafka
*/
func (h OpenEulerMeetingHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		var raw OpenEulerMeetingRaw
		msgBodyErr := json.Unmarshal(message.Value, &raw)
		if msgBodyErr != nil {
			return msgBodyErr
		}
		err := kfklib.Publish(config.OpenEulerMeetingConfigInstance.Publish, nil, message.Value)
		if err != nil {
			logrus.Error(err)
			return err
		}
		session.MarkMessage(message, "")
		logrus.Info("send meeting success")
	}
	return nil
}

type OpenEulerMeetingRaw struct {
	Action string `json:"action"`
	Msg    struct {
		Topic     string      `json:"topic"`
		Platform  interface{} `json:"mplatform"`
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
