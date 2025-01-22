package commands

import (
	"fmt"
	"log/slog"

	"github.com/bwmarrin/discordgo"
	"github.com/kotorkovsciy/discord-kiss-bot/pkg/logger"
)

type KissCommand struct{}

func (c *KissCommand) Name() string {
	return "kiss"
}

func (c *KissCommand) Description() string {
	return "Send a kiss to another user"
}

func (c *KissCommand) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionUser,
			Name:        "user",
			Description: "Who do you want to send a kiss to?",
			Required:    true,
		},
	}
}

func (c *KissCommand) Handle(s *discordgo.Session, i *discordgo.InteractionCreate) {
	log := logger.GetLogger()

	user := i.ApplicationCommandData().Options[0].UserValue(s)
	message := fmt.Sprintf("%s sent a kiss to %s! 💋", i.Member.User.Mention(), user.Mention())

	log.Info("Handling 'kiss' command",
		slog.String("sender", i.Member.User.Username),
		slog.String("recipient", user.Username))

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
		},
	})
	if err != nil {
		log.Error("Failed to respond to 'kiss' command",
			slog.String("sender", i.Member.User.Username),
			slog.String("recipient", user.Username),
			slog.String("error", err.Error()))
	} else {
		log.Info("Successfully sent a kiss message",
			slog.String("sender", i.Member.User.Username),
			slog.String("recipient", user.Username))
	}
}
