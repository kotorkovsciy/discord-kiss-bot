package bot

import (
	"log/slog"

	"github.com/bwmarrin/discordgo"
	"github.com/kotorkovsciy/discord-kiss-bot/internal/commands"
	"github.com/kotorkovsciy/discord-kiss-bot/pkg/config"
	"github.com/kotorkovsciy/discord-kiss-bot/pkg/logger"
)

type Bot struct {
	session *discordgo.Session
}

func New(cfg *config.Config) (*Bot, error) {
	log := logger.GetLogger()

	session, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		log.Error("Failed to create Discord session", slog.String("error", err.Error()))
		return nil, err
	}

	session.AddHandler(commands.HandleInteraction)

	err = session.Open()
	if err != nil {
		log.Error("Failed to open Discord session", slog.String("error", err.Error()))
		return nil, err
	}
	log.Info("Bot successfully started")

	err = commands.Register(session, cfg.GuildID)
	if err != nil {
		log.Error("Failed to register commands", slog.String("error", err.Error()))
		return nil, err
	}

	return &Bot{session: session}, nil
}

func (b *Bot) Close() {
	log := logger.GetLogger()

	if err := commands.Unregister(b.session); err != nil {
		log.Warn("Error while unregistering commands", slog.String("error", err.Error()))
	}
	if err := b.session.Close(); err != nil {
		log.Warn("Error while closing the session", slog.String("error", err.Error()))
	} else {
		log.Info("Bot session closed successfully")
	}
}
