package kafka

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	writer *kafka.Writer
}

func NewKafkaProducer(broker []string, topic string) *KafkaProducer {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(broker...),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	return &KafkaProducer{writer: writer}
}

func (p *KafkaProducer) Publish(message []byte) error {
	err := p.writer.WriteMessages(context.Background(), kafka.Message{
		Value: message,
	})
	if err != nil {
		return fmt.Errorf("failed to send message to Kafka: %w", err)
	}
	log.Printf("Message sent to Kafka: %s", message)
	return nil
}

func (p *KafkaProducer) Close() {
	if err := p.writer.Close(); err != nil {
		log.Fatal("Failed to close Kafka producer: ", err)
	}
}
