package client

import (
	"github.com/arslanovdi/omp-bot/internal/app/path"
	"github.com/arslanovdi/omp-bot/internal/model/user"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type ClientService interface {
	Describe(ClientID uint64) (*user.Client, error)
	List(cursor uint64, limit uint64) ([]user.Client, error)
	Get(cursor uint64) (user.Client, error)
	Create(user.Client) (uint64, error)
	Update(clientID uint64, client user.Client) error
	Remove(clientID uint64) (bool, error)
}

const limit = 10 // кол-во client выдаваемое за 1 раз

type clientCommander struct {
	bot           *tgbotapi.BotAPI
	clientService ClientService
}

func NewClientCommander(bot *tgbotapi.BotAPI, service ClientService) *clientCommander {

	return &clientCommander{
		bot:           bot,
		clientService: service,
	}
}

// HandleCallback перебор кнопок и вызов соттветствующего обработчика
func (c *clientCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("ClientCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

// HandleCommand перебор команд и вызов соттветствующего обработчика
func (c *clientCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "get":
		c.Get(msg)
	case "list":
		c.List(msg)
	case "delete":
		c.Delete(msg)
	case "new":
		c.New(msg)
	case "edit":
		c.Edit(msg)
	default:
		c.Default(msg)
	}
}

// errorResponseCommand возвращает сообщение об ошибке в бот
func (c *clientCommander) errorResponseCommand(message *tgbotapi.Message, resp string) {
	msg := tgbotapi.NewMessage(
		message.Chat.ID,
		resp,
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("ClientCommander.errorResponseCommand: error sending reply message to chat - %v", err)
	}
}

// errorResponseCallback возвращает сообщение об ошибке в бот
func (c *clientCommander) errorResponseCallback(callback *tgbotapi.CallbackQuery, resp string) {
	msg := tgbotapi.NewMessage(
		callback.Message.Chat.ID,
		resp,
	)
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("ClientCommander.errorResponseCallback: error sending reply message to chat - %v", err)
	}
}
