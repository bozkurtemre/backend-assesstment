package app

import (
	"log"
	"os"
	"os/signal"

	"github.com/bozkurtemre/backend-assesstment/src/worker/internal/config"
	"github.com/bozkurtemre/backend-assesstment/src/worker/internal/consumer"
	"github.com/bozkurtemre/backend-assesstment/src/worker/internal/database"
	"github.com/bozkurtemre/backend-assesstment/src/worker/internal/event"
	"github.com/bozkurtemre/backend-assesstment/src/worker/internal/logger"
	"github.com/bozkurtemre/backend-assesstment/src/worker/internal/wallet"
)

func Run() {
	loggerInstance, err := logger.NewLogger("events.log", "EVENT")
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}
	defer loggerInstance.Close()

	database.ConnectDB()

	walletRepository := wallet.NewWalletRepository(database.DB)
	walletService := wallet.NewWalletService(walletRepository)
	eventService := event.NewEventService(walletService, loggerInstance)

	pandaConsumer, err := consumer.NewPandaConsumer([]string{config.Config("BROKER_CONNSTR")}, eventService)
	if err != nil {
		log.Fatalf("failed to create consumer: %v", err)
	}

	err = pandaConsumer.ConsumeData("user-wallet-events")
	if err != nil {
		log.Fatalf("failed to consume data: %v", err)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	<-signals
}
