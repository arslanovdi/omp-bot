package _package

import (
	"github.com/arslanovdi/omp-bot/internal/app/path"
	"github.com/arslanovdi/omp-bot/internal/model/logistic"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type PackageService interface {
	Describe(PackageID uint64) (*logistic.Package, error)
	List(cursor uint64, limit uint64) ([]logistic.Package, error)
	Get(cursor uint64) (logistic.Package, error)
	Create(logistic.Package) (uint64, error)
	Update(packageID uint64, pkg logistic.Package) error
	Remove(packageID uint64) (bool, error)
}

const limit = 10 // кол-во package выдаваемое за 1 раз

type packageCommander struct {
	bot            *tgbotapi.BotAPI
	packageService PackageService
}

func NewPackageCommander(bot *tgbotapi.BotAPI, service PackageService) *packageCommander {

	return &packageCommander{
		bot:            bot,
		packageService: service,
	}
}

// HandleCallback перебор кнопок и вызов соттветствующего обработчика
func (c *packageCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("PackageCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

// HandleCommand перебор команд и вызов соттветствующего обработчика
func (c *packageCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
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
func (c *packageCommander) errorResponseCommand(message *tgbotapi.Message, resp string) {
	msg := tgbotapi.NewMessage(
		message.Chat.ID,
		resp,
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("PackageCommander.errorResponseCommand: error sending reply message to chat - %v", err)
	}
}

// errorResponseCallback возвращает сообщение об ошибке в бот
func (c *packageCommander) errorResponseCallback(callback *tgbotapi.CallbackQuery, resp string) {
	msg := tgbotapi.NewMessage(
		callback.Message.Chat.ID,
		resp,
	)
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("PackageCommander.errorResponseCallback: error sending reply message to chat - %v", err)
	}
}
