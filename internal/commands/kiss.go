package commands

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
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
	user := i.ApplicationCommandData().Options[0].UserValue(s)
	message := fmt.Sprintf("%s sent a kiss to %s! 💋", i.Member.User.Mention(), user.Mention())

	log.Printf("Handling 'kiss' command: %s sent a kiss to %s.", i.Member.User.Username, user.Username)

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
		},
	})
	if err != nil {
		log.Printf("Failed to respond to 'kiss' command: %v", err)
	} else {
		log.Printf("Successfully sent a kiss message from %s to %s.", i.Member.User.Username, user.Username)
	}
}
