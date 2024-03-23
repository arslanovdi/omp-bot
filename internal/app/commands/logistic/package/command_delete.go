package _package

import (
	"fmt"
	"log/slog"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *packageCommander) Delete(message *tgbotapi.Message) {

	log := slog.With("func", "packageCommander.Delete")

	args := message.CommandArguments()

	id, err := strconv.Atoi(args)
	if err != nil {
		c.errorResponseCommand(message, fmt.Sprintf("wrong args %v\n", args))
		log.Info("wrong args", slog.Any("args", args))
		return
	}

	ok, err := c.packageService.Delete(uint64(id))
	if err != nil {
		c.errorResponseCommand(message, fmt.Sprintf("Fail to delete package with id %d", id))
		log.Error("fail to delete package", slog.Uint64("id", uint64(id)), slog.Any("error", err))
		return
	}

	var msg tgbotapi.MessageConfig
	if ok {
		msg = tgbotapi.NewMessage(
			message.Chat.ID,
			fmt.Sprintf("Package with id: %d deleted", id),
		)
	} else {
		msg = tgbotapi.NewMessage(
			message.Chat.ID,
			fmt.Sprintf("Package with id: %d NOT deleted", id),
		)
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Error("error sending reply message to chat", slog.Any("error", err))
	}
}
