package main

import (
	"github.com/arslanovdi/omp-bot/internal/logger"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	routerPkg "github.com/arslanovdi/omp-bot/internal/app/router"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

const level = slog.LevelDebug // log level

func main() {
	logger.InitializeLogger(level) // slog logger

	_ = godotenv.Load()

	token, found := os.LookupEnv("TOKEN")
	if !found {
		slog.Warn("environment variable TOKEN not found in .env")
		os.Exit(1)
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		slog.Warn("Failed to create new bot", err)
		os.Exit(1)
	}

	// Uncomment if you want debugging
	// bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)

	routerHandler := routerPkg.NewRouter(bot) // Создаем обработчик телегрм бота

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT) // подписываем канал на сигналы завершения процесса
	for {
		select {
		case update := <-updates:
			routerHandler.HandleUpdate(update)
		case <-stop:
			slog.Info("Graceful shutdown")
			return
		}
	}

}
