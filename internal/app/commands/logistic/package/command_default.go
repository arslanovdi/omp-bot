package _package

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Default ответ при неверной команде
func (c *packageCommander) Default(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, "You wrote: "+message.Text)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("PackageCommander.Help: error sending reply message to chat - %v", err)
	}
}
