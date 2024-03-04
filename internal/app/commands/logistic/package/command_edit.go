package _package

import (
	"fmt"
	"github.com/arslanovdi/omp-bot/internal/model/logistic"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

// Edit обработка команды /edit бота
func (c *packageCommander) Edit(message *tgbotapi.Message) {
	args := message.CommandArguments()
	id := uint64(0)
	name := ""
	_, err := fmt.Sscanf(args, "%d %s", &id, &name)
	if err != nil || len(name) == 0 {
		c.errorResponseCommand(message, fmt.Sprintf("wrong args %v\n", args))
		log.Println("wrong args", args)
		return
	}

	err = c.packageService.Update(id, logistic.Package{Name: name})
	if err != nil {
		c.errorResponseCommand(message, fmt.Sprintf("Fail to edit package with id %d", id))
		log.Printf("fail to edit package with id %d: %v", id, err)
		return
	}

	msg := tgbotapi.NewMessage(
		message.Chat.ID,
		fmt.Sprintf("package id: %d renamed to %s", id, name),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("PackageCommander.New: error sending reply message to chat - %v", err)
	}
}
