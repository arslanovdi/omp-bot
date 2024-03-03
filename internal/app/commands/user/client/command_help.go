package client

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
)

func (c *clientCommander) Help(message *tgbotapi.Message) {
	str := strings.Builder{}
	str.WriteString("/help__user__client - help\n")
	str.WriteString("/get__user__client - get client (id)\n")
	str.WriteString("/list__user__client - list of clients\n")
	str.WriteString("/delete__user__client - delete client (id)\n")
	str.WriteString("/new__user__client - new client (name)\n")
	str.WriteString("/edit__user__client - set new client name (id, name)\n")

	msg := tgbotapi.NewMessage(message.Chat.ID,
		str.String(),
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("ClientCommander.Help: error sending reply message to chat - %v", err)
	}
}
