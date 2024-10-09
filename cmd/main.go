package main

import (
	"log"

	"github.com/dmarts05/forocoches-newsletter-codes/internal/codesextractor"
	"github.com/dmarts05/forocoches-newsletter-codes/internal/gmailclient"
	"github.com/dmarts05/forocoches-newsletter-codes/internal/telegramsender"
)

func main() {
	println(`
$$$$$$$$\  $$$$$$\         $$$$$$\   $$$$$$\  $$$$$$$\  $$$$$$$$\  $$$$$$\  
$$  _____|$$  __$$\       $$  __$$\ $$  __$$\ $$  __$$\ $$  _____|$$  __$$\ 
$$ |      $$ /  \__|      $$ /  \__|$$ /  $$ |$$ |  $$ |$$ |      $$ /  \__|
$$$$$\    $$ |            $$ |      $$ |  $$ |$$ |  $$ |$$$$$\    \$$$$$$\  
$$  __|   $$ |            $$ |      $$ |  $$ |$$ |  $$ |$$  __|    \____$$\ 
$$ |      $$ |  $$\       $$ |  $$\ $$ |  $$ |$$ |  $$ |$$ |      $$\   $$ |
$$ |      \$$$$$$  |      \$$$$$$  | $$$$$$  |$$$$$$$  |$$$$$$$$\ \$$$$$$  |
\__|       \______/        \______/  \______/ \_______/ \________| \______/
	`)

	gc, err := gmailclient.New("credentials.json", "token.json")
	if err != nil {
		log.Fatalf("Unable to create Gmail client: %v", err)
	}
	body, err := gc.GetLatestEmailFromSender("forocoches@substack.com")
	if err != nil {
		log.Fatalf("Unable to get email body: %v", err)
	}

	message := codesextractor.GenerateFCCodesMessageFromText(body)
	println(message)

	ts, err := telegramsender.New("telegram.json")
	if err != nil {
		log.Printf("Unable to create Telegram sender, no message will be sent: %v", err)
	} else {
		err = ts.SendMessage(message)
		if err != nil {
			log.Printf("Unable to send message via Telegram: %v", err)
		} else {
			log.Println("Message sent via Telegram successfully")
		}
	}

}
