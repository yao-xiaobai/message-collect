package plugin

import (
	"github.com/IBM/sarama"
	"github.com/opensourceways/message-collect/common/kafka"
	"github.com/opensourceways/message-collect/config"
	"github.com/sirupsen/logrus"
)

type EurBuildPlugin struct {
}

func (p EurBuildPlugin) StartConsume() {
	config.InitEurBuildConfig()
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

func (h EurGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		kafka.KfkProducer.SendMessage(config.EurBuildConfigInstance.Publish, message.Value)
		session.MarkMessage(message, "")
		logrus.Info("send eur build success")
	}
	return nil
}
