package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/atran25/hexaprogress/data"
	"github.com/atran25/hexaprogress/internal/command"
	"github.com/atran25/hexaprogress/internal/config"
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

func main() {
	if err := run(); err != nil {
		log.Err(err).Msg("init run failed")
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("getting config: %w", err)
	}
	log.Info().Str("DiscordToken", cfg.DiscordToken).Msg("Config loaded.")

	// discordgo bot setup & connection
	s, err := discordgo.New("Bot " + cfg.DiscordToken)
	if err != nil {
		return fmt.Errorf("creating discord session: %w", err)
	}
	defer s.Close()

	// Add command handlers
	classData, err := data.ReadClassData("data/class.json")
	if err != nil {
		return err
	}
	bossData, err := data.ReadBossData("data/boss.json")
	if err != nil {
		return err
	}
	hexaData, err := data.ReadHexaData("data/hexaLevelUpChart.json")
	if err != nil {
		return err
	}
	botCommands := command.NewBotCommands(classData, bossData, hexaData)
	botCommandsHandler, err := botCommands.GetAllCommandsHandler()
	if err != nil {
		return err
	}
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		handler, ok := botCommandsHandler[i.ApplicationCommandData().Name]
		if ok {
			handler(s, i)
		}
	})

	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Info().Str("username", s.State.User.Username).Str("discriminator", s.State.User.Discriminator).Msg("Bot is logged in.")
	})
	err = s.Open()
	if err != nil {
		return fmt.Errorf("opening connection to discord: %w", err)
	}

	allCommands := botCommands.GetAllCommands()
	syncCommands(s, cfg.GuildID, allCommands)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Info().Msg("Bot is now running.  Press CTRL-C to exit.")
	<-stop

	return nil
}

func syncCommands(s *discordgo.Session, guildID string, desiredCommandList []*discordgo.ApplicationCommand) {
	existingCommands, err := s.ApplicationCommands(s.State.User.ID, guildID)
	if err != nil {
		log.Fatal().Err(err).Str("guildID", guildID).Msg("Failed to fetch commands for guild")
		return
	}

	desiredMap := make(map[string]*discordgo.ApplicationCommand)
	for _, cmd := range desiredCommandList {
		desiredMap[cmd.Name] = cmd
	}

	existingMap := make(map[string]*discordgo.ApplicationCommand)
	for _, cmd := range existingCommands {
		existingMap[cmd.Name] = cmd
	}

	// Delete commands not in the desired list
	for _, cmd := range existingCommands {
		if _, found := desiredMap[cmd.Name]; !found {
			err := s.ApplicationCommandDelete(s.State.User.ID, guildID, cmd.ID)
			if err != nil {
				log.Info().Str("command", cmd.Name).Str("id", cmd.ID).Str("guildID", guildID).Msg("Failed to delete command")
			} else {
				log.Info().Str("command", cmd.Name).Str("id", cmd.ID).Str("guildID", guildID).Msg("Successfully deleted command")
			}
		}
	}

	// Create or update existing commands
	for _, cmd := range desiredCommandList {
		if existingCmd, found := existingMap[cmd.Name]; found {
			// Edit existing command
			_, err := s.ApplicationCommandEdit(s.State.User.ID, guildID, existingCmd.ID, cmd)
			if err != nil {
				log.Info().Str("command", cmd.Name).Str("id", cmd.ID).Str("guildID", guildID).Msg("Failed to edit command")
			} else {
				log.Info().Str("command", cmd.Name).Str("id", cmd.ID).Str("guildID", guildID).Msg("Successfully edited command")
			}
		} else {
			// Create new command
			_, err := s.ApplicationCommandCreate(s.State.User.ID, guildID, cmd)
			if err != nil {
				log.Info().Str("command", cmd.Name).Str("guildID", guildID).Msg("Failed to create command")
			} else {
				log.Info().Str("command", cmd.Name).Str("guildID", guildID).Msg("Successfully created command")
			}
		}
	}
}
