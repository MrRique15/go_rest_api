package kafka

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
)

var KafkaTopics = map[string]string{
	"order_sec": "order_sec",
	"inventory": "inventory",
	"payment": "payment",
	"shipping": "shipping",
}

var OrdersSEC_KafkaEvents = []string{
	"event.order.process",                    // 0

	"event.inventory.rollback_failed",        // 1
	"event.inventory.rollback_success",       // 2
	"event.inventory.reservation_failed",     // 3
	"event.inventory.reservation_succeeded",  // 4

	"event.payment.verify_failed",            // 5
	"event.payment.rollback_failed",          // 6
	"event.payment.rollback_success",         // 7
	"event.payment.verify_succeeded",         // 8

	"event.shipping.start_failed",            // 9
	"event.shipping.start_succeeded",         // 10
	"event.shipping.rollback_failed",         // 11
	"event.shipping.rollback_succeeded",      // 12
}

var Inventory_KafkaEvents = []string{
	"event.inventory.verify",
	"event.inventory.rollback",
}

var Payment_KafkaEvents = []string{
	"event.payment.verify",
	"event.payment.rollback",
}

var Shipping_KafkaEvents = []string{
	"event.shipping.start",
	"event.shipping.rollback",
}

func StartKafkaProducer(kafka_host string) (*sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 10
    config.Producer.Return.Successes = true

    producer, err := sarama.NewSyncProducer([]string{kafka_host}, config)
    if err != nil {
        log.Fatal("Failed to start Kafka producer:", err)
    }

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
	config.Consumer.Offsets.AutoCommit.Enable = true
	config.Consumer.Offsets.AutoCommit.Interval = 2 * time.Minute
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

    consumer, err := sarama.NewConsumer([]string{kafka_host}, config)
    if err != nil {
        log.Fatalf("Erro ao criar consumidor: %v", err)
    }

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