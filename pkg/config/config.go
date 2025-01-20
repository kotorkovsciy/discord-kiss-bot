package config

import (
	"errors"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Token   string
	GuildID string
}

func LoadConfig() (*Config, error) {
	if val, ok := os.LookupEnv("path_env"); ok {
		viper.SetConfigFile(val)
	} else {
		viper.SetConfigFile(".env")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	token := viper.GetString("DISCORD_BOT_TOKEN")
	if token == "" {
		return nil, errors.New("missing DISCORD_BOT_TOKEN in configuration")
	}

	return &Config{
		Token:   token,
		GuildID: viper.GetString("GuildID"),
	}, nil
}
