package gif

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/simpleCalculus/bot/gif/config"
	"github.com/simpleCalculus/bot/gif/types"
)

// Gif Метод который сравнивает сегодняшний курса рубля с вчерашним.
// Если курс по отношению к рублю за сегодня стал выше вчерашнего, то отдаем
// рандомную отсюда https://giphy.com/search/rich
// если ниже отсюда https://giphy.com/search/broke
func Gif() string {
	openExchange := config.NewOpenExchange()
	giphy := config.NewGiphy()

	today := getCurrency(openExchange.Latest())
	yesterday := getCurrency(openExchange.Yesterday())

	log.Printf("current dollar to ruble exchange rate: %f", today.Rates.RUB)
	log.Printf("yesterday's dollar to ruble exchange rate:: %f", yesterday.Rates.RUB)

	if today.Rates.RUB > yesterday.Rates.RUB {
		return getGif(giphy, "rich")
	} else {
		return getGif(giphy, "broke")
	}
}

// getCurrency метод который вернет структуру Latest
// в котором есть данные курса доллара к рублю
func getCurrency(url string) types.Latest {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	bytes, _ := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()

	var data types.Latest
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		log.Println(err)
	}
	return data
}

// getGif метод который вернет ссылку на рандомную gif в зависимости от tag
func getGif(giphy *config.Giphy, tag string) string {

	res, err := http.Get(giphy.Gif(tag))
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = res.Body.Close() }()

	var gif types.Gif
	_ = json.NewDecoder(res.Body).Decode(&gif)

	downloadGif(gif.Data.Images.DownsizedMedium.Url, gif.Data.Title)

	return fmt.Sprintf("%s", gif.Data.Images.DownsizedMedium.Url)
}

// downloadGif скачивает gif по адресу imgUrl и сохранять
// его в папку images с названием title
func downloadGif(imgUrl string, title string) {
	res, err := http.Get(imgUrl)
	if err != nil {
		log.Println("Error in http.Get(): " + err.Error())
		return
	}
	defer func() { _ = res.Body.Close() }()

	if res.StatusCode != 200 {
		log.Println("Received non 200 response code")
		return
	}

	gifName := "images/" + strings.ReplaceAll(title, " ", "_") + ".gif"

	file, err := os.Create(gifName)
	if err != nil {
		log.Printf("Error, can't create file: %s", err)
		return
	}
	defer func() { _ = file.Close() }()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		fmt.Printf("Error, can't copy bytes to file: %s", err)
	}
}
