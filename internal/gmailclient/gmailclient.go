package gmailclient

import (
	"encoding/base64"
	"errors"
	"log"
	"time"

	"google.golang.org/api/gmail/v1"
)

// Simplified Gmail client that uses the Gmail API.
type GmailClient struct {
	service *gmail.Service
}

// Creates a new Gmail client.
func New(credentialsFileName string, tokenFileName string) (*GmailClient, error) {
	service, err := getGmailService(credentialsFileName, tokenFileName)
	if err != nil {
		return nil, err
	}

	return &GmailClient{
		service,
	}, nil
}

// Obtains the body of the first email received today from a given sender.
func (gc *GmailClient) GetTodaysEmailFromSender(sender string) (string, error) {
	today := time.Now().Format("2006/01/02")
	tomorrow := time.Now().AddDate(0, 0, 1).Format("2006/01/02")
	query := "from:" + sender + " after:" + today + " before:" + tomorrow
	log.Printf("Querying Gmail with: %s", query)

	r, err := gc.service.Users.Messages.List("me").Q(query).Do()
	if err != nil {
		return "", err
	}
	if len(r.Messages) == 0 {
		return "", errors.New("no messages found")
	}

	message, err := gc.service.Users.Messages.Get("me", r.Messages[0].Id).Do()
	if err != nil {
		return "", err
	}

	body, err := extractEmailBody(message)
	if err != nil {
		return "", err
	}

	return body, nil
}

// Extracts the body of a Gmail message. It assumes that the message has only one part.
func extractEmailBody(message *gmail.Message) (string, error) {
	part := message.Payload.Parts[0]
	data, err := base64.URLEncoding.DecodeString(part.Body.Data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
