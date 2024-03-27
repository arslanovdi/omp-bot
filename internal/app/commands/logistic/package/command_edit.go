package _package

import (
	"errors"
	"fmt"
	"github.com/arslanovdi/omp-bot/internal/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	"strings"
	"time"
)

// Edit обработка команды /edit бота
func (c *packageCommander) Edit(message *tgbotapi.Message) {

	log := slog.With("func", "packageCommander.Edit")

	args := message.CommandArguments()

	pkg := model.Package{}

	var err error
	// Обработка опционального параметра Weight
	switch strings.Count(args, " ") {
	case 1:
		_, err = fmt.Sscanf(args, "%d %s", &pkg.ID, &pkg.Title)
		if err != nil || len(pkg.Title) == 0 {
			log.Info("wrong args", slog.Any("args", args), slog.String("error", err.Error()))
			err = fmt.Errorf("wrong args %v\n", args)
		}
	case 2:
		pkg.Weight = new(uint64)
		_, err = fmt.Sscanf(args, "%d %s %d", &pkg.ID, &pkg.Title, pkg.Weight)
		if err != nil || len(pkg.Title) == 0 || *pkg.Weight == 0 {
			log.Info("wrong args", slog.Any("args", args), slog.String("error", err.Error()))
			err = fmt.Errorf("wrong args %v\n", args)
		}
	default:
		log.Info("wrong args count", slog.Any("args", args))
		err = fmt.Errorf("wrong args %v\n", args)
	}

	if err != nil {
		c.errorResponseCommand(message, err.Error())
		return
	}

	pkg.Updated = new(time.Time)
	*pkg.Updated = time.Now()

	err = c.packageService.Update(pkg)
	if err != nil {
		log.Error("fail to edit package", slog.Uint64("id", pkg.ID), slog.String("error", err.Error()))
		if errors.Is(err, model.ErrNotFound) {
			c.errorResponseCommand(message, fmt.Sprintf("Package not found"))
			return
		}
		c.errorResponseCommand(message, fmt.Sprintf("Fail to edit package with id %d", pkg.ID))
		return
	}

	var msg tgbotapi.MessageConfig
	msg = tgbotapi.NewMessage(
		message.Chat.ID,
		fmt.Sprintf("Package with id: %d updated", pkg.ID),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Error("error sending reply message to chat", slog.String("error", err.Error()))
	}

	log.Debug("Package updated", slog.Any("package", pkg))
}
