package _package

import (
	"fmt"
	"github.com/arslanovdi/omp-bot/internal/model/logistic"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

// New обработка команды /new бота
func (c *packageCommander) New(message *tgbotapi.Message) {
	name := message.CommandArguments()

	if len(name) == 0 {
		c.errorResponseCommand(message, fmt.Sprintf("wrong args"))
		log.Printf("fail to create package with name %v", name)
		return
	}

	id, err := c.packageService.Create(logistic.Package{Title: name})

	if err != nil {
		c.errorResponseCommand(message, fmt.Sprintf("Fail to create package with name %v", name))
		log.Printf("fail to create package with name %v: %v", name, err)
		return
	}

	msg := tgbotapi.NewMessage(
		message.Chat.ID,
		fmt.Sprintf("package %v created with id: %d", name, id),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("PackageCommander.New: error sending reply message to chat - %v", err)
	}
}
