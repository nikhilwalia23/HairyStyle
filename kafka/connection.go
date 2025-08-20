// pkg/kafka/connection.go
package kafka

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

func NewKafkaWriter(brokers []string, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(brokers...),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
		Async:    false, // Make it true if you want async writes
	}
}

func NewKafkaReader(brokers []string, topic, groupID string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
		MaxWait:  time.Second,
	})
}

// PrepareMessage creates a Kafka message with given key and value.
// Controller only needs to call this without importing kafka-go.
func PrepareMessage(key, value string) kafka.Message {
	return kafka.Message{
		Key:   []byte(key),
		Value: []byte(value),
	}
}

// SendMessage sends a message to Kafka using a writer.
func SendMessage(ctx context.Context, writer *kafka.Writer, msg kafka.Message) error {
	return writer.WriteMessages(ctx, msg)
}
