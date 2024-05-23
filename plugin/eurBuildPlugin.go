package plugin

import (
	"github.com/IBM/sarama"
	"github.com/opensourceways/message-collect/common/kafka"
)

type EurBuildPlugin struct {
}

func (p EurBuildPlugin) StartConsume() {
	cfg := kafka.ConsumeConfig{
		Topic:   "test_message_center",
		Address: "182.160.6.195:9094",
		Group:   "message-collect",
		Offset:  sarama.OffsetNewest,
	}

	h := EurGroupHandler{}
	kafka.ConsumeGroup(cfg, &h)
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
		kafka.KfkProducer.SendMessage("eur_build_raw", message.Value)
	}
	return nil
}
