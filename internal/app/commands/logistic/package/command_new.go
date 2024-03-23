package _package

import (
	"fmt"
	"github.com/arslanovdi/omp-bot/internal/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	"time"
)

// New обработка команды /new бота
func (c *packageCommander) New(message *tgbotapi.Message) {

	log := slog.With("func", "packageCommander.New")

	name := message.CommandArguments()

	if len(name) == 0 {
		c.errorResponseCommand(message, fmt.Sprintf("wrong args"))
		log.Info("fail to create package", slog.String("name", name))
		return
	}

	id, err := c.packageService.Create(model.Package{
		Title:     name,
		Weight:    0, // TODO добавить вес в параметры команд телеграма
		CreatedAt: time.Now(),
	})

	if err != nil {
		c.errorResponseCommand(message, fmt.Sprintf("Fail to create package with name %v", name))
		log.Error("fail to create package", slog.String("name", name), slog.Any("error", err))
		return
	}

	msg := tgbotapi.NewMessage(
		message.Chat.ID,
		fmt.Sprintf("package %v created with id: %d", name, id),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Error("error sending reply message to chat", slog.Any("error", err))
	}
}
