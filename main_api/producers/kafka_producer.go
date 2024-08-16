package producers

import (
	"errors"

	"main_api/configs"

	"github.com/IBM/sarama"
)

type KafkaProducer struct {
	kafkaProducer *sarama.SyncProducer
}

func NewKafkaProducer() *KafkaProducer {
	producer, error := configs.ConnectKafka()

	if error != nil {
		panic(errors.New("error connecting to kafka"))
	}

	kafkaRepository := KafkaProducer{
		kafkaProducer: producer,
	}

	return &kafkaRepository
}

func (kr KafkaProducer) SendKafkaEvent(topic string, message string) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	_, _, err := (*kr.kafkaProducer).SendMessage(msg)

	if err != nil {
		return errors.New("error sending message to kafka")
	}

	return nil
}
