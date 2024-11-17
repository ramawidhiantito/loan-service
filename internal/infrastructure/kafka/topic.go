package kafka

import (
	"log"

	"github.com/segmentio/kafka-go"
)

func CreateTopic(brokers []string, topic string, partitions int, replicationFactor int) error {
	// Create a Kafka connection for the topic's administrative operations
	conn, err := kafka.Dial("tcp", brokers[0])
	if err != nil {
		return err
	}
	defer conn.Close()

	// Create topic configuration
	topicConfig := kafka.TopicConfig{
		Topic:             topic,
		NumPartitions:     partitions,
		ReplicationFactor: replicationFactor,
	}

	// Attempt to create the topic
	if err := conn.CreateTopics(topicConfig); err != nil {
		return err
	}

	log.Printf("Topic '%s' created successfully", topic)
	return nil
}
