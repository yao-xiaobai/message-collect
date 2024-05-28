package kafka

import (
	"github.com/IBM/sarama"
	"log"
)

var (
	KfkProducer *KafkaProducer
)

// KafkaProducer 封装了 consume 生产者的功能
type KafkaProducer struct {
	producer sarama.SyncProducer
}

func Init(config *Config) {
	brokers := []string{config.Address}
	KfkProducer = NewKafkaProducer(brokers)
}

type Config struct {
	Address        string `json:"address" required:"true"`
	Version        string `json:"version"` // e.g 2.1.0
	MQCert         string `json:"mq_cert"`
	Username       string `json:"user_name"`
	Password       string `json:"password"`
	Algorithm      string `json:"algorithm"`
	SkipCertVerify bool   `json:"skip_cert_verify"`
}

// NewKafkaProducer 创建一个新的 consume 生产者
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
func (kp *KafkaProducer) SendMessage(topic string, message []byte) error {
	// 创建要发送的消息
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(message),
	}

	// 发送消息并处理结果
	_, _, err := kp.producer.SendMessage(msg)
	if err != nil {
		return err
	}
	return nil
}

// Close 关闭 consume 生产者连接
func (kp *KafkaProducer) Close() {
	if err := kp.producer.Close(); err != nil {
		log.Printf("Error closing consume producer: %v\n", err)
	}
}
