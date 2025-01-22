package commands

import (
	"log/slog"

	"github.com/bwmarrin/discordgo"
	"github.com/kotorkovsciy/discord-kiss-bot/pkg/logger"
)

var commandHandlers = []CommandHandler{
	&KissCommand{},
}

func Register(s *discordgo.Session, guildID string) error {
	log := logger.GetLogger()

	for _, cmd := range commandHandlers {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, guildID, &discordgo.ApplicationCommand{
			Name:        cmd.Name(),
			Description: cmd.Description(),
			Options:     cmd.Options(),
		})
		if err != nil {
			log.Error("Failed to register command",
				slog.String("command", cmd.Name()),
				slog.String("error", err.Error()))
			return err
		}

		if guildID != "" {
			log.Info("Command successfully registered for guild",
				slog.String("command", cmd.Name()),
				slog.String("guildID", guildID))
		} else {
			log.Info("Command successfully registered globally",
				slog.String("command", cmd.Name()))
		}
	}
	return nil
}

func Unregister(s *discordgo.Session) error {
	log := logger.GetLogger()

	cmds, err := s.ApplicationCommands(s.State.User.ID, "")
	if err != nil {
		log.Error("Failed to fetch commands for unregistration",
			slog.String("error", err.Error()))
		return err
	}

	for _, cmd := range cmds {
		err := s.ApplicationCommandDelete(s.State.User.ID, "", cmd.ID)
		if err != nil {
			log.Warn("Failed to delete command",
				slog.String("command", cmd.Name),
				slog.String("error", err.Error()))
		} else {
			log.Info("Command successfully deleted",
				slog.String("command", cmd.Name))
		}
	}
	return nil
}

func HandleInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	log := logger.GetLogger()

	for _, cmd := range commandHandlers {
		if i.ApplicationCommandData().Name == cmd.Name() {
			log.Info("Handling interaction for command",
				slog.String("command", cmd.Name()))
			cmd.Handle(s, i)
			return
		}
	}
	log.Warn("Received interaction with unknown command",
		slog.String("command", i.ApplicationCommandData().Name))
}
