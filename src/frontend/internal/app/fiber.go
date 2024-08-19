package app

import (
	"log"
	"os"

	"github.com/bozkurtemre/backend-assesstment/src/frontend/internal/router"
	"github.com/bozkurtemre/backend-assesstment/src/frontend/pkg/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	_ "github.com/joho/godotenv"
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

	port := ":" + os.Getenv("PORT")
	log.Fatal(app.Listen(port))
}
