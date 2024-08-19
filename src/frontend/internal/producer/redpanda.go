package producer

import (
	"github.com/IBM/sarama"
)

type PandaProducer struct {
	topic    string
	producer sarama.SyncProducer
}

func NewPandaProducer(topic string, brokers []string) (*PandaProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &PandaProducer{topic: topic, producer: producer}, nil
}

func (r *PandaProducer) SendData(data []byte) error {
	message := &sarama.ProducerMessage{
		Topic: r.topic,
		Value: sarama.ByteEncoder(data),
	}
	_, _, err := r.producer.SendMessage(message)

	return err
}
