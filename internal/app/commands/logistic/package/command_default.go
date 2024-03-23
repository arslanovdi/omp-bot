package _package

import (
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Default ответ при неверной команде
func (c *packageCommander) Default(message *tgbotapi.Message) {

	log := slog.With("func", "packageCommander.Default")

	log.Info("wrong command", slog.String("username", message.From.UserName), slog.String("message", message.Text))

	msg := tgbotapi.NewMessage(message.Chat.ID, "You wrote: "+message.Text)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Error("error sending reply message to chat", slog.Any("error", err))
	}
}
