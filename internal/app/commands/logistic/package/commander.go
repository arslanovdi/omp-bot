package _package

import (
	"github.com/arslanovdi/omp-bot/internal/app/path"
	"github.com/arslanovdi/omp-bot/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
)

const limit = 10 // кол-во package выдаваемое за 1 раз

type packageCommander struct {
	bot            *tgbotapi.BotAPI
	packageService *service.LogisticPackageService
}

func NewPackageCommander(bot *tgbotapi.BotAPI, service *service.LogisticPackageService) *packageCommander {

	return &packageCommander{
		bot:            bot,
		packageService: service,
	}
}

// HandleCallback перебор кнопок и вызов соттветствующего обработчика
func (c *packageCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {

	log := slog.With("func", "PackageCommander.HandleCallback")

	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Info("unknown callback name", slog.String("callback name", callbackPath.CallbackName))
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

	log := slog.With("func", "PackageCommander.errorResponseCommand")

	msg := tgbotapi.NewMessage(
		message.Chat.ID,
		resp,
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Error("error sending reply message to chat", slog.Any("error", err))
	}
}

// errorResponseCallback возвращает сообщение об ошибке в бот
func (c *packageCommander) errorResponseCallback(callback *tgbotapi.CallbackQuery, resp string) {

	log := slog.With("func", "packageCommander.errorResponseCallback")

	msg := tgbotapi.NewMessage(
		callback.Message.Chat.ID,
		resp,
	)
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Error("error sending reply message to chat", slog.Any("error", err))
	}
}
