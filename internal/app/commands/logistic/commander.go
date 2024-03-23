package logistic

import (
	"github.com/arslanovdi/omp-bot/internal/app/commands/logistic/package"
	"github.com/arslanovdi/omp-bot/internal/service"
	"log/slog"

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
func NewLogisticCommander(bot *tgbotapi.BotAPI, pkgService *service.LogisticPackageService) *logisticCommander {
	return &logisticCommander{
		bot:              bot,
		packageCommander: _package.NewPackageCommander(bot, pkgService),
	}
}

func (c *logisticCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {

	log := slog.With("func", "logisticCommander.HandleCallback")

	switch callbackPath.Subdomain {
	case "package":
		c.packageCommander.HandleCallback(callback, callbackPath)
	default:
		log.Info("unknown subdomain", slog.String("subdomain", callbackPath.Subdomain))
	}
}

func (c *logisticCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {

	log := slog.With("func", "logisticCommander.HandleCommand")

	switch commandPath.Subdomain {
	case "package":
		c.packageCommander.HandleCommand(msg, commandPath)
	default:
		log.Info("unknown subdomain", slog.String("subdomain", commandPath.Subdomain))
	}
}
