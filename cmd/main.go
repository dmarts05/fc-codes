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
	found := run()

	// If the wait flag is set (non-zero), run the program in a loop
	if *waitTime > 0 {
		for {
			if found {
				break
			}
			log.Printf("Waiting for %d seconds before the next execution...\n", *waitTime)
			time.Sleep(time.Duration(*waitTime) * time.Second)
			found = run()
		}
	}
}

func run() bool {
	// Create Gmail client
	log.Println("Creating Gmail client...")
	gc, err := gmailclient.New("credentials.json", "token.json")
	if err != nil {
		log.Printf("Unable to create Gmail client: %v", err)
		return false
	}

	// Get latest email body
	log.Println("Getting email body...")
	body, err := gc.GetTodaysEmailFromSender("forocoches@substack.com")
	if err != nil {
		log.Printf("Unable to get email body: %v", err)
		return false
	}

	// Extract message
	log.Println("Extracting FC codes...")
	message := codesextractor.GenerateFCCodesMessageFromText(body)
	fmt.Println(message)

	// Create Telegram sender and send the message
	log.Println("Sending message to Telegram...")
	ts, err := telegramsender.New("telegram.json")
	if err != nil {
		log.Printf("Unable to create Telegram sender: %v", err)
		return false
	}

	err = ts.SendMessage(message)
	if err != nil {
		log.Printf("Unable to send message: %v", err)
		return false
	}

	log.Println("Message sent successfully!")
	return true
}
