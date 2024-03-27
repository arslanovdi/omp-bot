package _package

import (
	"errors"
	"fmt"
	"github.com/arslanovdi/omp-bot/internal/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
)

func (c *packageCommander) Get(message *tgbotapi.Message) {

	log := slog.With("func", "packageCommander.Get")

	args := message.CommandArguments()

	id := uint64(0)
	_, err := fmt.Sscanf(args, "%d", &id)
	if err != nil {
		c.errorResponseCommand(message, fmt.Sprintf("wrong args %v", args))
		log.Info("wrong args", slog.Any("args", args), slog.String("error", err.Error()))
		return
	}

	pkg, err := c.packageService.Get(id)
	if err != nil {
		log.Error("fail to get product", slog.Uint64("id", id), slog.String("error", err.Error()))
		if errors.Is(err, model.ErrNotFound) {
			c.errorResponseCommand(message, fmt.Sprintf("Package with id: %d not found.\n", id))
			return
		}
		c.errorResponseCommand(message, "fail to get product\n")
		return
	}

	msg := tgbotapi.NewMessage(
		message.Chat.ID,
		pkg.String(),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Error("error sending reply message to chat", slog.String("error", err.Error()))
	}

	log.Debug("get package", slog.Any("pkg", pkg))
}
