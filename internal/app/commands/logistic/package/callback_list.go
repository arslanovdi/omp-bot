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

type CallbackListData struct {
	Offset int `json:"offset"`
}

// CallbackList обработка реакции на нажатие кнопки
func (c *packageCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {

	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		c.errorResponseCallback(callback, fmt.Sprintf("внутренняя ошибка"))
		log.Printf("packageCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}

	outputMsgText := strings.Builder{}
	outputMsgText.WriteString("These are our packages: \n\n")

	packages, err := c.packageService.List(uint64(parsedData.Offset)+limit, limit) // Запрашиваем клиентов со смещением

	var endOfList bool

	if err != nil {
		if errors.Is(err, logistic.EndOfList) {
			endOfList = true
		} else {
			c.errorResponseCallback(callback, fmt.Sprintf("внутренняя ошибка"))
			log.Printf("Ошибка получения списка: %v", err)
			return
		}
	}

	for _, p := range packages {
		outputMsgText.WriteString(p.Name)
		outputMsgText.WriteString("\n")
	}

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, outputMsgText.String())

	if !endOfList {
		serializedData, _ := json.Marshal(CallbackListData{ // данные сериализуемые в кнопку
			Offset: parsedData.Offset + limit,
		})
		callbackPath.CallbackData = string(serializedData)

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup( // добавляем кнопку в ответ
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
			),
		)
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("PackageCommander.CallbackList: error sending reply message to chat - %v", err)
	}
}
