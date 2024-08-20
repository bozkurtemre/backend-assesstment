package event

import (
	"encoding/json"
	"log"
	"sort"
	"sync"

	"github.com/bozkurtemre/backend-assesstment/src/frontend/internal/producer"
)

type Service struct {
	pandaProducer *producer.PandaProducer
}

type EventService interface {
	AddEvents(events Events, wg *sync.WaitGroup)
}

func NewEventService(pandaProducer *producer.PandaProducer) *Service {
	return &Service{pandaProducer: pandaProducer}
}

func (s *Service) AddEvents(events Events, wg *sync.WaitGroup) {
	defer wg.Done()

	sort.Slice(events.Events, func(i, j int) bool {
		eventI := events.Events[i].Time
		eventJ := events.Events[j].Time

		return eventI.After(eventJ)
	})

	for _, event := range events.Events {
		wg.Add(1)
		go func(e Event) {
			defer wg.Done()

			if err := ValidateEvent(e); err != nil {
				log.Printf("failed to validate event: %v", err)
			}

			eventData, err := json.Marshal(e)
			if err != nil {
				log.Printf("failed to marshal event data: %v", err)
			}

			if err = s.pandaProducer.SendData(eventData); err != nil {
				log.Printf("failed to process event data: %v", err)
			}
		}(event)
	}
}
