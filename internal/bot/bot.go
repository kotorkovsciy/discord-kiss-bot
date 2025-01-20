package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/kotorkovsciy/discord-kiss-bot/internal/commands"
	"github.com/kotorkovsciy/discord-kiss-bot/pkg/config"
)

type Bot struct {
	session *discordgo.Session
}

func New(cfg *config.Config) (*Bot, error) {
	session, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		return nil, err
	}

	session.AddHandler(commands.HandleInteraction)

	err = session.Open()
	if err != nil {
		return nil, err
	}
	log.Println("Bot successfully started.")

	err = commands.Register(session, cfg.GuildID)
	if err != nil {
		return nil, err
	}

	return &Bot{session: session}, nil
}

func (b *Bot) Close() {
	if err := commands.Unregister(b.session); err != nil {
		log.Printf("Error while unregistering commands: %v", err)
	}
	if err := b.session.Close(); err != nil {
		log.Printf("Error while closing the session: %v", err)
	} else {
		log.Println("Bot session closed successfully.")
	}
}
