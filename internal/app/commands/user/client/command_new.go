package client

import (
	"fmt"
	"github.com/arslanovdi/omp-bot/internal/model/user"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

// New обработка команды /new бота
func (c *clientCommander) New(message *tgbotapi.Message) {
	name := message.CommandArguments()

	if len(name) == 0 {
		c.errorResponseCommand(message, fmt.Sprintf("wrong args"))
		log.Printf("fail to create client with name %v", name)
		return
	}

	id, err := c.clientService.Create(user.Client{Name: name})

	if err != nil {
		c.errorResponseCommand(message, fmt.Sprintf("Fail to create client with name %v", name))
		log.Printf("fail to create client with name %v: %v", name, err)
		return
	}

	msg := tgbotapi.NewMessage(
		message.Chat.ID,
		fmt.Sprintf("client %v created with id: %d", name, id),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("ClientCommander.New: error sending reply message to chat - %v", err)
	}
}
