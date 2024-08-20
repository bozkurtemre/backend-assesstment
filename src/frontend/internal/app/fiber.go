package app

import (
	"log"

	"github.com/bozkurtemre/backend-assesstment/src/frontend/internal/config"
	"github.com/bozkurtemre/backend-assesstment/src/frontend/internal/database"
	"github.com/bozkurtemre/backend-assesstment/src/frontend/internal/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func Run() {
	app := fiber.New(fiber.Config{
		Prefork: false,
	})

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(requestid.New())

	database.ConnectDB()
	router.SetupRoutes(app)

	port := ":" + config.Config("PORT")
	log.Fatal(app.Listen(port))
}
