package packaging

import (
	"encoding/json"
	"errors"
	"github.com/arslanovdi/omp-bot/internal/model"
	"log/slog"
	"strings"

	"github.com/arslanovdi/omp-bot/internal/app/path"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// List обработка команды /list бота
func (c *Commander) List(message *tgbotapi.Message) {

	log := slog.With("func", "Commander.List")

	outputMsgText := strings.Builder{}
	outputMsgText.WriteString("These are all our packages: \n\n")

	packages, err := c.packageService.List(1, limit)
	var endOfList bool

	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			c.errorResponseCommand(message, "packages not found")
			return
		}
		if errors.Is(err, model.ErrEndOfList) {
			endOfList = true
		} else {
			c.errorResponseCommand(message, "Ошибка получения списка")
			log.Error("fail to get list of packages", slog.String("error", err.Error()))
			return
		}
	}

	for _, p := range packages {
		outputMsgText.WriteString(p.String())
		outputMsgText.WriteString("\n")
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, outputMsgText.String())

	if !endOfList {
		serializedData, _ := json.Marshal(CallbackListData{ // данные сериализуемые в кнопку
			Offset: 1,
		})

		callbackPath := path.CallbackPath{ // собираем структуру кнопки
			Domain:       "logistic",
			Subdomain:    "package",
			CallbackName: "list",
			CallbackData: string(serializedData),
		}

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup( // добавляем кнопку в ответ
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
			),
		)
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Error("error sending reply message to chat", slog.String("error", err.Error()))
	}

	log.Debug("Command List packages", slog.Uint64("offset", 1), slog.Uint64("limit", limit))
}
