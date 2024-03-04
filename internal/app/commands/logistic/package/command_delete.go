package _package

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *packageCommander) Delete(message *tgbotapi.Message) {
	args := message.CommandArguments()

	id, err := strconv.Atoi(args)
	if err != nil {
		c.errorResponseCommand(message, fmt.Sprintf("wrong args %v\n", args))
		log.Println("wrong args", args)
		return
	}

	ok, err := c.packageService.Remove(uint64(id))
	if err != nil {
		c.errorResponseCommand(message, fmt.Sprintf("Fail to delete package with id %d", id))
		log.Printf("fail to delete package with id %d: %v", id, err)
		return
	}

	var msg tgbotapi.MessageConfig
	if ok {
		msg = tgbotapi.NewMessage(
			message.Chat.ID,
			fmt.Sprintf("Package with id: %d deleted", id),
		)
	} else {
		msg = tgbotapi.NewMessage(
			message.Chat.ID,
			fmt.Sprintf("Package with id: %d NOT deleted", id),
		)
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("PackageCommander.Delete: error sending reply message to chat - %v", err)
	}
}
