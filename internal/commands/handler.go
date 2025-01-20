package commands

import "github.com/bwmarrin/discordgo"

type CommandHandler interface {
	Name() string
	Description() string
	Options() []*discordgo.ApplicationCommandOption
	Handle(s *discordgo.Session, i *discordgo.InteractionCreate)
}
