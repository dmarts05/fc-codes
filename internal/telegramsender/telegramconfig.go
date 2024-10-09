package telegramsender

import (
	"encoding/json"
	"os"
)

// Required configuration to send messages via Telegram.
type telegramConfig struct {
	Token  string `json:"token"`
	ChatID string `json:"chat_id"`
}

// Reads the Telegram configuration from the given file.
func getTelegramConfig(telegramFileName string) (telegramConfig, error) {
	file, err := os.Open(telegramFileName)
	if err != nil {
		return telegramConfig{}, err
	}
	defer file.Close()

	var config telegramConfig
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return telegramConfig{}, err
	}

	return config, nil
}
