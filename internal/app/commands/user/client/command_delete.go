package client

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *clientCommander) Delete(message *tgbotapi.Message) {
	args := message.CommandArguments()

	id, err := strconv.Atoi(args)
	if err != nil {
		c.errorResponseCommand(message, fmt.Sprintf("wrong args %v\n", args))
		log.Println("wrong args", args)
		return
	}

	ok, err := c.clientService.Remove(uint64(id))
	if err != nil {
		c.errorResponseCommand(message, fmt.Sprintf("Fail to delete client with id %d", id))
		log.Printf("fail to delete client with id %d: %v", id, err)
		return
	}

	var msg tgbotapi.MessageConfig
	if ok {
		msg = tgbotapi.NewMessage(
			message.Chat.ID,
			fmt.Sprintf("Client with id: %d deleted", id),
		)
	} else {
		msg = tgbotapi.NewMessage(
			message.Chat.ID,
			fmt.Sprintf("Client with id: %d NOT deleted", id),
		)
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("ClientCommander.Delete: error sending reply message to chat - %v", err)
	}
}
