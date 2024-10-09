package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/dmarts05/fc-codes/internal/codesextractor"
	"github.com/dmarts05/fc-codes/internal/gmailclient"
	"github.com/dmarts05/fc-codes/internal/telegramsender"
)

func main() {
	// Define flags
	waitTime := flag.Int("wait", 0, "Wait time in seconds between executions (0 to run once)")
	flag.Parse()

	// Show title
	fmt.Println(`
$$$$$$$$\  $$$$$$\         $$$$$$\   $$$$$$\  $$$$$$$\  $$$$$$$$\  $$$$$$\  
$$  _____|$$  __$$\       $$  __$$\ $$  __$$\ $$  __$$\ $$  _____|$$  __$$\ 
$$ |      $$ /  \__|      $$ /  \__|$$ /  $$ |$$ |  $$ |$$ |      $$ /  \__|
$$$$$\    $$ |            $$ |      $$ |  $$ |$$ |  $$ |$$$$$\    \$$$$$$\  
$$  __|   $$ |            $$ |      $$ |  $$ |$$ |  $$ |$$  __|    \____$$\ 
$$ |      $$ |  $$\       $$ |  $$\ $$ |  $$ |$$ |  $$ |$$ |      $$\   $$ |
$$ |      \$$$$$$  |      \$$$$$$  | $$$$$$  |$$$$$$$  |$$$$$$$$\ \$$$$$$  |
\__|       \______/        \______/  \______/ \_______/ \________| \______/
	`)

	// Run once before the loop so even if the wait flag is set to 0, the program runs at least once
	run()

	// If the wait flag is set (non-zero), run the program in a loop
	if *waitTime > 0 {
		for {
			fmt.Printf("Waiting for %d seconds before the next execution...\n", *waitTime)
			time.Sleep(time.Duration(*waitTime) * time.Second)
			run()
		}
	}
}

func run() {
	// Create Gmail client
	gc, err := gmailclient.New("credentials.json", "token.json")
	if err != nil {
		log.Fatalf("Unable to create Gmail client: %v", err)
	}

	// Get latest email body
	body, err := gc.GetLatestEmailFromSender("forocoches@substack.com")
	if err != nil {
		log.Fatalf("Unable to get email body: %v", err)
	}

	// Extract message
	message := codesextractor.GenerateFCCodesMessageFromText(body)
	fmt.Println(message)

	// Create Telegram sender and send the message
	ts, err := telegramsender.New("telegram.json")
	if err != nil {
		log.Fatalf("Unable to create Telegram sender: %v", err)
	}

	err = ts.SendMessage(message)
	if err != nil {
		log.Fatalf("Unable to send message: %v", err)
	}
}
