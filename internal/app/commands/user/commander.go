package user

import (
	"github.com/arslanovdi/omp-bot/internal/app/commands/user/client"
	service "github.com/arslanovdi/omp-bot/internal/service/user/client"
	"log"

	"github.com/arslanovdi/omp-bot/internal/app/path"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ClientCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)
	New(inputMsg *tgbotapi.Message)  // return error not implemented
	Edit(inputMsg *tgbotapi.Message) // return error not implemented
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type userCommander struct {
	bot             *tgbotapi.BotAPI
	clientCommander ClientCommander
}

// NewUserCommander конструктор
func NewUserCommander(
	bot *tgbotapi.BotAPI,
) *userCommander {

	return &userCommander{
		bot: bot,
		// subdomainCommander
		clientCommander: client.NewClientCommander(bot, service.NewClientService()),
	}
}

func (c *userCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "client":
		c.clientCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("UserCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *userCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "client":
		c.clientCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("UserCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
