package _package

import (
	"fmt"
	"github.com/arslanovdi/omp-bot/internal/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	"time"
)

// Edit обработка команды /edit бота
func (c *packageCommander) Edit(message *tgbotapi.Message) {

	log := slog.With("func", "packageCommander.Edit")

	args := message.CommandArguments()

	pkg := model.Package{}

	_, err := fmt.Sscanf(args, "%d %s %d", &pkg.ID, &pkg.Title, &pkg.Weight)
	if err != nil || len(pkg.Title) == 0 || pkg.Weight <= 0 {
		c.errorResponseCommand(message, fmt.Sprintf("wrong args %v\n", args))
		log.Info("wrong args", slog.Any("args", args), slog.String("error", err.Error()))
		return
	}

	pkg.CreatedAt = time.Now()
	ok, err := c.packageService.Update(pkg.ID, pkg)
	if err != nil {
		c.errorResponseCommand(message, fmt.Sprintf("Fail to edit package with id %d", pkg.ID))
		log.Error("fail to edit package", slog.Uint64("id", pkg.ID), slog.String("error", err.Error()))
		return
	}

	var msg tgbotapi.MessageConfig
	if ok {
		msg = tgbotapi.NewMessage(
			message.Chat.ID,
			fmt.Sprintf("Package with id: %d updated", pkg.ID),
		)
		log.Debug("Package updated", slog.Uint64("id", pkg.ID))
	} else {
		msg = tgbotapi.NewMessage(
			message.Chat.ID,
			fmt.Sprintf("Package with id: %d NOT found", pkg.ID),
		)
		log.Debug("Package not found", slog.Uint64("id", pkg.ID))
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Error("error sending reply message to chat", slog.String("error", err.Error()))
	}
}
