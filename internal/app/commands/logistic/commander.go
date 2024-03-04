package logistic

import (
	"github.com/arslanovdi/omp-bot/internal/app/commands/logistic/package"
	service "github.com/arslanovdi/omp-bot/internal/service/logistic/package"
	"log"

	"github.com/arslanovdi/omp-bot/internal/app/path"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type PackageCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)
	New(inputMsg *tgbotapi.Message)  // return error not implemented
	Edit(inputMsg *tgbotapi.Message) // return error not implemented
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type logisticCommander struct {
	bot              *tgbotapi.BotAPI
	packageCommander PackageCommander
}

// NewLogisticCommander конструктор
func NewLogisticCommander(
	bot *tgbotapi.BotAPI,
) *logisticCommander {

	return &logisticCommander{
		bot: bot,
		// subdomainCommander
		packageCommander: _package.NewPackageCommander(bot, service.NewPackageService()),
	}
}

func (c *logisticCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "package":
		c.packageCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("LogisticCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *logisticCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "package":
		c.packageCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("LogisticCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
