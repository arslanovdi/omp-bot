package _package

import (
	"fmt"
	"github.com/arslanovdi/omp-bot/internal/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
)

// Edit обработка команды /edit бота
func (c *packageCommander) Edit(message *tgbotapi.Message) {

	log := slog.With("func", "packageCommander.Edit")

	args := message.CommandArguments()
	id := uint64(0)
	name := ""
	_, err := fmt.Sscanf(args, "%d %s", &id, &name)
	if err != nil || len(name) == 0 {
		c.errorResponseCommand(message, fmt.Sprintf("wrong args %v\n", args))
		log.Info("wrong args", slog.Any("args", args))
		return
	}

	err = c.packageService.Update(id, model.Package{Title: name})
	if err != nil {
		c.errorResponseCommand(message, fmt.Sprintf("Fail to edit package with id %d", id))
		log.Error("fail to edit package", slog.Uint64("id", id), slog.Any("error", err))
		return
	}

	msg := tgbotapi.NewMessage(
		message.Chat.ID,
		fmt.Sprintf("package id: %d renamed to %s", id, name),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Error("error sending reply message to chat", slog.Any("error", err))
	}
}
