package router

import (
	"log"
	"runtime/debug"

	"github.com/arslanovdi/omp-bot/internal/app/commands/user"
	"github.com/arslanovdi/omp-bot/internal/app/path"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(callback *tgbotapi.Message, commandPath path.CommandPath)
}

type Router struct {
	// bot
	bot *tgbotapi.BotAPI

	// demoCommander
	userCommander Commander // экземпляр интерфейса обрабатывающий сообщения телеграм бота
	// user
	// access
	// buy
	// delivery
	// recommendation
	// travel
	// loyalty
	// bank
	// subscription
	// license
	// insurance
	// payment
	// storage
	// streaming
	// business
	// work
	// service
	// exchange
	// estate
	// rating
	// security
	// cinema
	// logistic
	// product
	// education
}

func NewRouter(
	bot *tgbotapi.BotAPI,
) *Router {
	return &Router{
		// bot
		bot: bot,
		// demoCommander
		userCommander: user.NewUserCommander(bot),
		//demo.NewDemoCommander(bot), // приведение структуры к интерфейсу
		// user
		// access
		// buy
		// delivery
		// recommendation
		// travel
		// loyalty
		// bank
		// subscription
		// license
		// insurance
		// payment
		// storage
		// streaming
		// business
		// work
		// service
		// exchange
		// estate
		// rating
		// security
		// cinema
		// logistic
		// product
		// education
	}
}

// HandleUpdate обработка сообщений телеграм бота
func (c *Router) HandleUpdate(update tgbotapi.Update) {
	defer func() { // рековер исключений, чтобы бот не умер
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered from panic: %v\n%v", panicValue, string(debug.Stack()))
		}
	}()

	switch {
	case update.CallbackQuery != nil:
		c.handleCallback(update.CallbackQuery) // обработка кнопок
	case update.Message != nil:
		c.handleMessage(update.Message) // обработка сообщений
	}
}

// handleCallback обработка нажатия кнопок
func (c *Router) handleCallback(callback *tgbotapi.CallbackQuery) {
	callbackPath, err := path.ParseCallback(callback.Data)
	if err != nil {
		log.Printf("Router.handleCallback: error parsing callback data `%s` - %v", callback.Data, err)
		return
	}

	switch callbackPath.Domain {
	case "demo":
		break //c.demoCommander.HandleCallback(callback, callbackPath)
	case "user":
		c.userCommander.HandleCallback(callback, callbackPath)
	case "access":
		break
	case "buy":
		break
	case "delivery":
		break
	case "recommendation":
		break
	case "travel":
		break
	case "loyalty":
		break
	case "bank":
		break
	case "subscription":
		break
	case "license":
		break
	case "insurance":
		break
	case "payment":
		break
	case "storage":
		break
	case "streaming":
		break
	case "business":
		break
	case "work":
		break
	case "service":
		break
	case "exchange":
		break
	case "estate":
		break
	case "rating":
		break
	case "security":
		break
	case "cinema":
		break
	case "logistic":
		break
	case "product":
		break
	case "education":
		break
	default:
		log.Printf("Router.handleCallback: unknown domain - %s", callbackPath.Domain)
	}
}

// handleMessage обработка команд
func (c *Router) handleMessage(msg *tgbotapi.Message) {
	if !msg.IsCommand() {
		c.showCommandFormat(msg)

		return
	}

	commandPath, err := path.ParseCommand(msg.Command())
	if err != nil {
		log.Printf("Router.handleCallback: error parsing callback data `%s` - %v", msg.Command(), err)
		return
	}

	switch commandPath.Domain {
	case "demo":
		break //c.demoCommander.HandleCommand(msg, commandPath)
	case "user":
		c.userCommander.HandleCommand(msg, commandPath)
	case "access":
		break
	case "buy":
		break
	case "delivery":
		break
	case "recommendation":
		break
	case "travel":
		break
	case "loyalty":
		break
	case "bank":
		break
	case "subscription":
		break
	case "license":
		break
	case "insurance":
		break
	case "payment":
		break
	case "storage":
		break
	case "streaming":
		break
	case "business":
		break
	case "work":
		break
	case "service":
		break
	case "exchange":
		break
	case "estate":
		break
	case "rating":
		break
	case "security":
		break
	case "cinema":
		break
	case "logistic":
		break
	case "product":
		break
	case "education":
		break
	default:
		log.Printf("Router.handleCallback: unknown domain - %s", commandPath.Domain)
	}
}

// showCommandFormat выдача в бот сообщения с форматом команд
func (c *Router) showCommandFormat(inputMessage *tgbotapi.Message) {
	outputMsg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Command format: /{command}__{domain}__{subdomain}")

	_, err := c.bot.Send(outputMsg)
	if err != nil {
		log.Printf("Router.showCommandFormat: error sending reply message to chat - %v", err)
	}
}
