package _package

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/arslanovdi/omp-bot/internal/model/logistic"
	"log"
	"strings"

	"github.com/arslanovdi/omp-bot/internal/app/path"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *packageCommander) List(message *tgbotapi.Message) {
	outputMsgText := strings.Builder{}
	outputMsgText.WriteString("These are all our packages: \n\n")

	packages, err := c.packageService.List(1, limit)
	var endOfList bool

	if err != nil {
		if errors.Is(err, logistic.EndOfList) {
			endOfList = true
		} else {
			c.errorResponseCommand(message, fmt.Sprintf("Ошибка получения списка"))
			log.Printf("Ошибка получения списка: %v", err)
			return
		}
	}

	for _, p := range packages {
		outputMsgText.WriteString(p.Title)
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
		log.Printf("PackageCommander.List: error sending reply message to chat - %v", err)
	}
}
