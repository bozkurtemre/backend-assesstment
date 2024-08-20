package consumer

import (
	"log"

	"github.com/bozkurtemre/backend-assesstment/src/worker/internal/event"

	"github.com/IBM/sarama"
)

type PandaConsumer struct {
	consumer     sarama.Consumer
	eventService event.EventService
}

func NewPandaConsumer(brokers []string, eventService event.EventService) (*PandaConsumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &PandaConsumer{consumer: consumer, eventService: eventService}, nil
}

func (c *PandaConsumer) ConsumeData(topic string) error {
	partitions, err := c.consumer.Partitions(topic)
	if err != nil {
		return err
	}

	for _, partition := range partitions {
		part, err := c.consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
		if err != nil {
			return err
		}

		go func(part sarama.PartitionConsumer) {
			defer part.Close()

			for message := range part.Messages() {
				err = c.eventService.ProcessEvent(message.Value)
				if err != nil {
					log.Printf("failed to process event: %v", err)
					continue
				}

				log.Printf("event processed successfully")
			}
		}(part)
	}

	return nil
}
