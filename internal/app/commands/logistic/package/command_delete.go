package _package

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
)

func (c *packageCommander) Delete(message *tgbotapi.Message) {

	log := slog.With("func", "packageCommander.Delete")

	args := message.CommandArguments()

	id := uint64(0)
	_, err := fmt.Sscanf(args, "%d", &id)
	if err != nil {
		c.errorResponseCommand(message, fmt.Sprintf("wrong args %v\n", args))
		log.Info("wrong args", slog.Any("args", args), slog.String("error", err.Error()))
		return
	}

	ok, err := c.packageService.Delete(id)
	if err != nil {
		c.errorResponseCommand(message, fmt.Sprintf("Fail to delete package with id %d", id))
		log.Error("fail to delete package", slog.Uint64("id", id), slog.String("error", err.Error()))
		return
	}

	var msg tgbotapi.MessageConfig
	if ok {
		msg = tgbotapi.NewMessage(
			message.Chat.ID,
			fmt.Sprintf("Package with id: %d deleted", id),
		)
		log.Info("Package deleted", slog.Uint64("id", uint64(id)))
	} else {
		msg = tgbotapi.NewMessage(
			message.Chat.ID,
			fmt.Sprintf("Package with id: %d NOT found", id),
		)
		log.Info("Package not found", slog.Uint64("id", uint64(id)))
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Error("error sending reply message to chat", slog.String("error", err.Error()))
	}
}
