package event

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/bozkurtemre/backend-assesstment/src/frontend/internal/producer"
)

type Service struct {
	pandaProducer *producer.PandaProducer
}

type EventService interface {
	ProcessEvents(events Events, wg *sync.WaitGroup)
}

func NewEventService(pandaProducer *producer.PandaProducer) *Service {
	return &Service{pandaProducer: pandaProducer}
}

func (s *Service) ProcessEvents(events Events, wg *sync.WaitGroup) {
	defer wg.Done()

	for _, event := range events.Events {
		wg.Add(1)
		go func(e Event) {
			defer wg.Done()
			if err := ValidateEvent(e); err != nil {
				log.Printf("Failed to validate event: %v", err)
			}

			eventData, err := json.Marshal(e)
			if err != nil {
				log.Printf("Failed to marshal event data: %v", err)
			}

			if err = s.pandaProducer.SendData(eventData); err != nil {
				log.Printf("Failed to process event data: %v", err)
			}
		}(event)
	}
}
