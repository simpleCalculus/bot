package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	TelegramBotToken string `json:"TelegramBotToken"`
	Host             string `json:"host"`
}

func New() Config {
	file, err := os.Open("config/config.json")
	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Panic(err)
	}
	return configuration
}
