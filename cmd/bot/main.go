package main

import (
	"context"
	"errors"
	routerPkg "github.com/arslanovdi/omp-bot/internal/app/router"
	"github.com/arslanovdi/omp-bot/internal/config"
	"github.com/arslanovdi/omp-bot/internal/grpc"
	"github.com/arslanovdi/omp-bot/internal/logger"
	"github.com/arslanovdi/omp-bot/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const level = slog.LevelDebug // log level

func main() {
	logger.InitializeLogger(level) // slog logger
	log := slog.With("func", "main")

	startCtx, cancel := context.WithTimeout(context.Background(), time.Minute) // контекст запуска приложения
	defer cancel()
	go func() {
		<-startCtx.Done()
		if errors.Is(startCtx.Err(), context.DeadlineExceeded) { // приложение зависло при запуске
			log.Warn("Application startup time exceeded")
			os.Exit(1)
		}
	}()

	err := config.ReadConfigYML("config.yml")
	if err != nil {
		log.Warn("Failed to read config", slog.String("error", err.Error()))
		os.Exit(1)
	}

	_ = godotenv.Load()

	token, found := os.LookupEnv("TOKEN")
	if !found {
		log.Warn("environment variable TOKEN not found in .env")
		os.Exit(1)
	}

	grpcClient := grpc.NewGrpcClient()
	packageService := service.NewPackageService(grpcClient)

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Warn("Failed to create new bot", slog.String("error", err.Error()))
		os.Exit(1)
	}

	log.Info("Telegram bot authorized on account ", slog.String("account", bot.Self.UserName))

	// Uncomment if you want debugging
	// bot.Debug = true

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u) // получаем канал обновлений телеграм бота

	routerHandler := routerPkg.NewRouter(bot, packageService) // Создаем обработчик телегрм бота

	cancel() // отменяем контекст запуска приложения
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT) // подписываем канал на сигналы завершения процесса
	for {
		select {
		case update := <-updates:
			routerHandler.HandleUpdate(update)
		case <-stop:
			slog.Info("Graceful shutdown")
			grpcClient.Close()
			return
		}
	}

}
