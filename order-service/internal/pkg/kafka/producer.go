package kafka

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/IBM/sarama"
)

type Producer[T any] struct {
	producer sarama.SyncProducer
	topic    string
}

func NewProducer[T any](cfg Config, topic string) (*Producer[T], error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll

	producer, err := sarama.NewSyncProducer(cfg.Hosts, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create producer: %w", err)
	}

	return &Producer[T]{
		producer: producer,
		topic:    topic,
	}, nil
}

func (p *Producer[T]) Produce(_ context.Context, value T) error {
	valueJSON, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("produce kafka msg marshall err: %w", err)
	}

	msg := &sarama.ProducerMessage{
		Topic: p.topic,
		Value: sarama.ByteEncoder(valueJSON),
	}

	_, _, err = p.producer.SendMessage(msg)
	if err != nil {
		return fmt.Errorf("produce kafka msg send err: %w", err)
	}

	return nil
}
