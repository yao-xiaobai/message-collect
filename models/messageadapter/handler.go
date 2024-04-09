package messageadapter

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/opensourceways/kafka-lib/mq"
	"message-collect/models/domain/event"
)

// Handler
type Handler interface {
	handle(message []byte) error
}

type EurHandler struct{}

func (eurHandler *EurHandler) handle(message []byte) error {
	fmt.Println(message)
	return nil
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
		var msg mq.Message
		err := json.Unmarshal(message.Value, &msg)
		if err != nil {
			return err
		}
		var eurBuild event.EurBuild

		err1 := json.Unmarshal(msg.Body, &eurBuild)
		if err1 != nil {
			return err
		}
		err3 := SendMsg(&eurBuild)
		if err3 != nil {
			return err3
		}
	}

	return nil
}
