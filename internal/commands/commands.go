package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var commandHandlers = []CommandHandler{
	&KissCommand{},
}

func Register(s *discordgo.Session, guildID string) error {
	for _, cmd := range commandHandlers {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, guildID, &discordgo.ApplicationCommand{
			Name:        cmd.Name(),
			Description: cmd.Description(),
			Options:     cmd.Options(),
		})
		if err != nil {
			return err
		}

		if guildID != "" {
			log.Printf("Command '%s' successfully registered for guild '%s'.", cmd.Name(), guildID)
		} else {
			log.Printf("Command '%s' successfully registered globally.", cmd.Name())
		}
	}
	return nil
}

func Unregister(s *discordgo.Session) error {
	cmds, err := s.ApplicationCommands(s.State.User.ID, "")
	if err != nil {
		return err
	}

	for _, cmd := range cmds {
		err := s.ApplicationCommandDelete(s.State.User.ID, "", cmd.ID)
		if err != nil {
			log.Printf("Failed to delete command '%s': %v", cmd.Name, err)
		} else {
			log.Printf("Command '%s' successfully deleted.", cmd.Name)
		}
	}
	return nil
}

func HandleInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	for _, cmd := range commandHandlers {
		if i.ApplicationCommandData().Name == cmd.Name() {
			log.Printf("Handling interaction for command '%s'.", cmd.Name())
			cmd.Handle(s, i)
			return
		}
	}
	log.Printf("Received interaction with unknown command '%s'.", i.ApplicationCommandData().Name)
}
