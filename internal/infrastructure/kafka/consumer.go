package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	reader *kafka.Reader
}

func NewKafkaConsumer(brokers []string, topic string, groupID string) *KafkaConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: groupID,
	})

	return &KafkaConsumer{reader: reader}
}

func (c *KafkaConsumer) ConsumeMessages() {
	for {
		msg, err := c.reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("Error while consuming message: %v", err)
		}
		log.Printf("Message received from Kafka: %s", string(msg.Value))

		// Handle the message here (e.g., process the loan event)
	}
}

func (c *KafkaConsumer) Close() {
	if err := c.reader.Close(); err != nil {
		log.Fatal("Failed to close Kafka consumer: ", err)
	}
}
