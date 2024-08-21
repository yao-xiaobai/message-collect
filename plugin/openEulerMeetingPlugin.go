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
