package _package

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
)

func (c *packageCommander) Help(message *tgbotapi.Message) {
	str := strings.Builder{}
	str.WriteString("/help__logistic__package - help\n")
	str.WriteString("/get__logistic__package - get package (id)\n")
	str.WriteString("/list__logistic__package - list of packages\n")
	str.WriteString("/delete__logistic__package - delete package (id)\n")
	str.WriteString("/new__logistic__package - new package (name)\n")
	str.WriteString("/edit__logistic__package - set new package name (id, name)\n")

	msg := tgbotapi.NewMessage(message.Chat.ID,
		str.String(),
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("PackageCommander.Help: error sending reply message to chat - %v", err)
	}
}
