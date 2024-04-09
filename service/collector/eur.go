package collector

import (
	"github.com/IBM/sarama"
	"message-collect/models/messageadapter"
)

func Consume() {
	cfg := messageadapter.ConsumeConfig{
		Topic:   "eur_build_raw",
		Address: "0.0.0.0:9092",
		Group:   "ssp_test",
		Offset:  sarama.OffsetNewest,
	}

	h := messageadapter.EurGroupHandler{}
	messageadapter.ConsumeGroup(cfg, &h)
}
