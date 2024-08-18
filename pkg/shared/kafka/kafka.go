package kafka

import (
	"errors"
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

func StartKafkaProducer(kafka_host string) (*sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 10
    config.Producer.Return.Successes = true

    producer, err := sarama.NewSyncProducer([]string{kafka_host}, config)
    if err != nil {
        log.Fatal("Failed to start Kafka producer:", err)
    }
    // defer producer.Close()

	if err != nil {
		log.Fatal("Failed to connect to Kafka:", err)
		return nil, err
	}

	fmt.Println("Started Kafka Producer")

	return &producer, nil
}

func StartKafkaConsumer(kafka_host string) (sarama.Consumer, error) {
    config := sarama.NewConfig()
    config.Consumer.Return.Errors = true

    consumer, err := sarama.NewConsumer([]string{kafka_host}, config)
    if err != nil {
        log.Fatalf("Erro ao criar consumidor: %v", err)
    }
    // defer consumer.Close()

	fmt.Println("Started Kafka Consumer")

	return consumer, nil
}

// -------------------------------------
// KafkaProducer struct
// -------------------------------------

type KafkaProducer struct {
	kafkaProducer *sarama.SyncProducer
}

func (kr *KafkaProducer) SendKafkaEvent(topic string, message string) error {
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

func NewKafkaProducer(kafka_host string) *KafkaProducer {
	producer, error := StartKafkaProducer(kafka_host)

	if error != nil {
		panic(errors.New("error connecting to kafka"))
	}

	kafkaProd := KafkaProducer{
		kafkaProducer: producer,
	}

	return &kafkaProd
}