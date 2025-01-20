package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kotorkovsciy/discord-kiss-bot/internal/bot"
	"github.com/kotorkovsciy/discord-kiss-bot/pkg/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	b, err := bot.New(cfg)
	if err != nil {
		log.Fatalf("Bot initialization failed: %v", err)
	}
	defer b.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	log.Println("Bot has been stopped gracefully.")
}
