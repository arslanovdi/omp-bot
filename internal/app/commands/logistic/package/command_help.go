package _package

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	"strings"
)

func (c *packageCommander) Help(message *tgbotapi.Message) {

	log := slog.With("func", "packageCommander.Help")

	str := strings.Builder{}
	str.WriteString("/help__logistic__package - help\n")
	str.WriteString("/get__logistic__package - get package (id)\n")
	str.WriteString("/list__logistic__package - list of packages\n")
	str.WriteString("/delete__logistic__package - delete package (id)\n")
	str.WriteString("/new__logistic__package - new package (name, weight)\n")
	str.WriteString("/edit__logistic__package - set new package name (id, name, weight)\n")

	msg := tgbotapi.NewMessage(message.Chat.ID,
		str.String(),
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Error("error sending reply message to chat", slog.Any("error", err))
	}
}
