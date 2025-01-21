package config

import (
	"errors"
	"os"
)

type Config struct {
	Token   string
	GuildID string
}

func LoadConfig() (*Config, error) {
	token := os.Getenv("DISCORD_BOT_TOKEN")

	if token == "" {
		return nil, errors.New("missing DISCORD_BOT_TOKEN in environment variables")
	}

	guildID := os.Getenv("GUILD_ID")

	return &Config{
		Token:   token,
		GuildID: guildID,
	}, nil
}
