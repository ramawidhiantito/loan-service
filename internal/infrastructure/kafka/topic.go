package kafka

import (
	"log"

	"github.com/segmentio/kafka-go"
)

func CreateTopic(brokers []string, topic string, partitions int, replicationFactor int) error {
	conn, err := kafka.Dial("tcp", brokers[0])
	if err != nil {
		return err
	}
	defer conn.Close()

	topicConfig := kafka.TopicConfig{
		Topic:             topic,
		NumPartitions:     partitions,
		ReplicationFactor: replicationFactor,
	}

	if err := conn.CreateTopics(topicConfig); err != nil {
		return err
	}

	log.Printf("Topic '%s' created successfully", topic)
	return nil
}
