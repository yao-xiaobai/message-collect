package kafka

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"log"
)

var (
	KfkProducer *KafkaProducer
)

// KafkaProducer 封装了 Kafka 生产者的功能
type KafkaProducer struct {
	producer sarama.SyncProducer
}

func Init() {
	brokers := []string{"127.0.0.1:9092"}
	KfkProducer = NewKafkaProducer(brokers)
}

// NewKafkaProducer 创建一个新的 Kafka 生产者
func NewKafkaProducer(brokerList []string) *KafkaProducer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		return nil
	}

	return &KafkaProducer{
		producer: producer,
	}
}

// SendMessage 发送消息到指定的主题
func (kp *KafkaProducer) SendMessage(topic string, message interface{}) error {
	// 创建要发送的消息
	jsonBytes, err := json.Marshal(message)
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(jsonBytes),
	}

	// 发送消息并处理结果
	_, _, err = kp.producer.SendMessage(msg)
	if err != nil {
		return err
	}
	return nil
}

// Close 关闭 Kafka 生产者连接
func (kp *KafkaProducer) Close() {
	if err := kp.producer.Close(); err != nil {
		log.Printf("Error closing Kafka producer: %v\n", err)
	}
}
