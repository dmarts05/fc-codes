package telegramsender

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

// Represents a sender that can send messages via Telegram.
type TelegramSender struct {
	config telegramConfig
}

// Creates a new TelegramSender with the provided configuration.
func New(telegramFileName string) (*TelegramSender, error) {
	config, err := getTelegramConfig(telegramFileName)
	if err != nil {
		return nil, err
	}

	return &TelegramSender{
		config,
	}, nil
}

// Sends a message via Telegram.
func (ts *TelegramSender) SendMessage(message string) error {
	sendUrl := ts.getBotUrl() + "/sendMessage"
	body, err := json.Marshal(map[string]string{
		"chat_id": ts.config.ChatID,
		"text":    message,
	})
	if err != nil {
		return err
	}

	r, err := http.Post(sendUrl, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	// Handle possible errors
	if r.StatusCode != http.StatusOK {
		return errors.New("failed to send message")
	}

	return nil
}

func (ts *TelegramSender) getBotUrl() string {
	return "https://api.telegram.org/bot" + ts.config.Token
}
