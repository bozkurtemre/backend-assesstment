package router

import (
	"fmt"
	"log"

	"github.com/bozkurtemre/backend-assesstment/src/frontend/internal/config"
	"github.com/bozkurtemre/backend-assesstment/src/frontend/internal/database"
	"github.com/bozkurtemre/backend-assesstment/src/frontend/internal/event"
	"github.com/bozkurtemre/backend-assesstment/src/frontend/internal/producer"
	"github.com/bozkurtemre/backend-assesstment/src/frontend/internal/wallet"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	walletRepository := wallet.NewWalletRepository(database.DB)
	walletService := wallet.NewWalletService(walletRepository)
	wallet.NewWalletHandler(app.Group("/"), walletService)

	pandaProducer, err := producer.NewPandaProducer("user-wallet-events", []string{config.Config("BROKER_CONNSTR")})
	if err != nil {
		log.Fatalf("failed to create producer: %v", err)
	}

	eventService := event.NewEventService(pandaProducer)
	event.NewEventHandler(app.Group("/"), eventService)

	app.All("*", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": fmt.Sprintf("Route '%s' does not exist in this API!", c.OriginalURL()),
		})
	})
}
