package router

import (
	"fmt"
	"log"
	"os"

	"github.com/bozkurtemre/backend-assesstment/src/frontend/internal/event"
	"github.com/bozkurtemre/backend-assesstment/src/frontend/internal/producer"
	"github.com/bozkurtemre/backend-assesstment/src/frontend/internal/wallet"
	"github.com/bozkurtemre/backend-assesstment/src/frontend/pkg/database"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Wallet
	walletRepository := wallet.NewWalletRepository(database.DB)
	walletService := wallet.NewWalletService(walletRepository)
	wallet.NewWalletHandler(app.Group("/"), walletService)

	// Event
	pandaProducer, err := producer.NewPandaProducer("wallet-events", []string{os.Getenv("BROKER_CONNSTR")})
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
	}

	eventService := event.NewEventService(pandaProducer)
	event.NewEventHandler(app.Group("/"), eventService)

	app.All("*", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "fail",
			"message": fmt.Sprintf("Route '%s' does not exist in this API!", c.OriginalURL()),
		})
	})
}
