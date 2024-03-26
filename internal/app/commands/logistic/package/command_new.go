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

	args := message.CommandArguments()

	pkg := model.Package{}

	_, err := fmt.Sscanf(args, "%s %d", &pkg.Title, &pkg.Weight)

	if len(pkg.Title) == 0 || pkg.Weight <= 0 || err != nil {
		c.errorResponseCommand(message, fmt.Sprintf("wrong args"))
		log.Info("wrong args", slog.String("package", pkg.String()))
		return
	}

	pkg.CreatedAt = time.Now()

	id, err := c.packageService.Create(pkg)

	if err != nil {
		c.errorResponseCommand(message, fmt.Sprintf("Fail to create package with title %v", pkg.Title))
		log.Error("fail to create package", slog.String("package", pkg.String()), slog.String("error", err.Error()))
		return
	}

	log.Debug("Package created", slog.Uint64("id", id), slog.String("package", pkg.String()))

	msg := tgbotapi.NewMessage(
		message.Chat.ID,
		fmt.Sprintf("package %v created with id: %d", pkg.Title, id),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Error("error sending reply message to chat", slog.String("error", err.Error()))
	}
}
