package plugin

import (
	"fmt"
	"github.com/IBM/sarama"
	"github.com/opensourceways/message-collect/common/kafka"
	"github.com/opensourceways/message-collect/utils"
)

type EurBuildPlugin struct {
}

func (p EurBuildPlugin) StartConsume() {
	cfg := new(kafka.ConsumeConfig)
	if err := utils.LoadFromYaml("config/eur_build_conf.yaml", cfg); err != nil {
		fmt.Println("Config初始化失败, err:", err)
		return
	}
	h := EurGroupHandler{}
	kafka.ConsumeGroup(*cfg, &h)
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
