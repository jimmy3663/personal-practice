package order

import (
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func streamhandler(message string) {
	producer, topic := KafkaProducer()

	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, nil)

}
