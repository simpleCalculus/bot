package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/simpleCalculus/bot/config"
	"github.com/simpleCalculus/bot/gif"
	"github.com/simpleCalculus/bot/lib/e"
	"github.com/simpleCalculus/bot/types"
)

const defaultMsg = "Доступные команды help и currency"

const help = `Бот проверяет курс доллар к рублю. 
Если курс по отношению к рублю за сегодня стал выше вчерашнего, то отдает рандомную отсюда https://giphy.com/search/rich. 
Иначе отдает отсюда https://giphy.com/search/broke.
Чтобы получить гиф введите currency`

func main() {
	c := config.New()

	botUrl := basePath(c.Host, c.TelegramBotToken)
	offset := 0

	for {
		updates, err := updates(botUrl, offset)
		if err != nil {
			log.Println("Smth went wrong: ", err.Error())
		}

		for _, update := range updates {
			switch update.Message.Text {
			case "currency":
				// отправляем ссылку на рандовную gif в зависимости от курса доллара к рублю
				update.Message.Text = gif.Gif()
				err = respondText(botUrl, update)
				offset = update.UpdateID + 1
			case "help":
				// отправляем инструкцию по пользованию
				update.Message.Text = help
				err = respondText(botUrl, update)
				offset = update.UpdateID + 1
			default:
				update.Message.Text = defaultMsg
				err = respondText(botUrl, update)
				offset = update.UpdateID + 1
			}
		}
	}
}

func basePath(host string, token string) string {
	return "https://" + host + "/bot" + token
}

func updates(botUrl string, offset int) ([]types.Update, error) {
	resp, err := http.Get(botUrl + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, e.Wrap("can't get request", err)
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, e.Wrap("can't read bytes", err)
	}

	var restResponse types.RestResponse
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, e.Wrap("can't unmarshal body", err)
	}

	return restResponse.Result, nil
}

func respondText(botUrl string, update types.Update) error {
	var botMessage types.BotMessage
	botMessage.ChatID = update.Message.Chat.ChatID
	botMessage.Text = update.Message.Text

	buf, err := json.Marshal(botMessage)
	if err != nil {
		return e.Wrap("can't marshal botMessage", err)
	}
	_, err = http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return e.Wrap("can't send message", err)
	}

	return nil
}
