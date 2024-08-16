package consumers

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

func OrdersConsumer(consumer sarama.Consumer) {
	// Escolhe o tópico e partição
	topic := "order"
	partition := int32(0)

	// Inicia a leitura das mensagens
	partitionConsumer, err := consumer.ConsumePartition(topic, partition, sarama.OffsetOldest)
	if err != nil {
		log.Fatalf("Erro ao consumir partição: %v", err)
	}
	defer partitionConsumer.Close()

	// Loop para escutar as mensagens
	for msg := range partitionConsumer.Messages() {
		fmt.Printf("Recebido: %s\n", string(msg.Value))
	}

	// Escuta erros (opcional)
	go func() {
		for err := range partitionConsumer.Errors() {
			log.Printf("Erro consumidor: %v", err)
		}
	}()

	// Mantenha a aplicação rodando
	select {}
}