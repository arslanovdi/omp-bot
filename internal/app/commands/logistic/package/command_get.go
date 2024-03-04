package _package

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *packageCommander) Get(message *tgbotapi.Message) {
	args := message.CommandArguments()

	id, err := strconv.Atoi(args)
	if err != nil {
		c.errorResponseCommand(message, fmt.Sprintf("wrong args %v", args))
		log.Println("wrong args", args)
		return
	}

	pkg, err := c.packageService.Get(uint64(id))
	if err != nil {
		c.errorResponseCommand(message, fmt.Sprintf("Package with id: %d not found.\n", id))
		log.Printf("fail to get product with id %d: %v", id, err)
		return
	}

	msg := tgbotapi.NewMessage(
		message.Chat.ID,
		pkg.String(),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("PackageCommander.Get: error sending reply message to chat - %v", err)
	}
}
