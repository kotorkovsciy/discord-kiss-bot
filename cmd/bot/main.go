package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/kotorkovsciy/discord-kiss-bot/internal/bot"
	"github.com/kotorkovsciy/discord-kiss-bot/pkg/config"
	"github.com/kotorkovsciy/discord-kiss-bot/pkg/logger"
)

func main() {
	log := logger.GetLogger()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Error("Failed to load configuration", slog.String("error", err.Error()))
	}

	b, err := bot.New(cfg)
	if err != nil {
		log.Error("Bot initialization failed", slog.String("error", err.Error()))
	}
	defer b.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	log.Info("Bot has been stopped gracefully.")
}
