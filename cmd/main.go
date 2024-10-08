package main

import (
	"log"

	"github.com/dmarts05/forocoches-newsletter-codes/internal/codesextractor"
	"github.com/dmarts05/forocoches-newsletter-codes/internal/gmailclient"
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

	gc := gmailclient.New("credentials.json", "token.json")
	body, err := gc.GetLatestEmailFromSender("forocoches@substack.com")
	if err != nil {
		log.Fatalf("Unable to get email body: %v", err)
	}

	codes := codesextractor.GetForocochesCodesFromText(body)
	for _, code := range codes {
		println(code)
	}
}
