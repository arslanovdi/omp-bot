package _package

import (
	"fmt"
	"log/slog"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *packageCommander) Get(message *tgbotapi.Message) {

	log := slog.With("func", "packageCommander.Get")

	args := message.CommandArguments()

	id, err := strconv.Atoi(args)
	if err != nil {
		c.errorResponseCommand(message, fmt.Sprintf("wrong args %v", args))
		log.Info("wrong args", slog.Any("args", args))
		return
	}

	pkg, err := c.packageService.Get(uint64(id))
	if err != nil {
		c.errorResponseCommand(message, fmt.Sprintf("Package with id: %d not found.\n", id))
		log.Error("fail to get product", slog.Uint64("id", uint64(id)), slog.Any("error", err))
		return
	}

	msg := tgbotapi.NewMessage(
		message.Chat.ID,
		pkg.String(),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Error("error sending reply message to chat", slog.Any("error", err))
	}
}
