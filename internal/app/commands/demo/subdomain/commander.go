package subdomain

import (
	"log"

	"github.com/arslanovdi/omp-bot/internal/app/path"
	"github.com/arslanovdi/omp-bot/internal/service/demo/subdomain"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type DemoSubdomainCommander struct {
	bot              *tgbotapi.BotAPI
	subdomainService *subdomain.Service
}

// NewDemoSubdomainCommander конструктор
func NewDemoSubdomainCommander(
	bot *tgbotapi.BotAPI,
) *DemoSubdomainCommander {
	subdomainService := subdomain.NewService()

	return &DemoSubdomainCommander{
		bot:              bot,
		subdomainService: subdomainService,
	}
}

// HandleCallback перебор кнопок и вызов соттветствующего обработчика
func (c *DemoSubdomainCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("DemoSubdomainCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

// HandleCommand перебор команд и вызов соттветствующего обработчика
func (c *DemoSubdomainCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	default:
		c.Default(msg)
	}
}
