package event

import (
	"sync"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	eventService EventService
}

func NewEventHandler(eventRoute fiber.Router, eventService EventService) {
	handler := &Handler{eventService}

	eventRoute.Post("/", handler.ProcessEvents)
}

func (h *Handler) ProcessEvents(c *fiber.Ctx) error {
	var eventsRequest Events

	if err := c.BodyParser(&eventsRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": err.Error()})
	}

	if err := Validator.Struct(eventsRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": err.Error()})
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go h.eventService.ProcessEvents(eventsRequest, &wg)
	wg.Wait()

	return c.SendStatus(fiber.StatusOK)
}
