package client

import (
	"fmt"
	"github.com/arslanovdi/omp-bot/internal/model/user"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

// Edit обработка команды /edit бота
func (c *clientCommander) Edit(message *tgbotapi.Message) {
	args := message.CommandArguments()
	id := uint64(0)
	name := ""
	_, err := fmt.Sscanf(args, "%d %s", &id, &name)
	if err != nil || len(name) == 0 {
		c.errorResponseCommand(message, fmt.Sprintf("wrong args %v\n", args))
		log.Println("wrong args", args)
		return
	}

	err = c.clientService.Update(id, user.Client{Name: name})
	if err != nil {
		c.errorResponseCommand(message, fmt.Sprintf("Fail to edit client with id %d", id))
		log.Printf("fail to edit client with id %d: %v", id, err)
		return
	}

	msg := tgbotapi.NewMessage(
		message.Chat.ID,
		fmt.Sprintf("client id: %d renamed to %s", id, name),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("ClientCommander.New: error sending reply message to chat - %v", err)
	}
}
