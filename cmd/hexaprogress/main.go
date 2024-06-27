package main

import (
	"fmt"
	"os"
	"os/signal"

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
	cfg, err := GetConfig()
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
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Info().Str("username", s.State.User.Username).Str("discriminator", s.State.User.Discriminator).Msg("Bot is logged in.")
	})
	err = s.Open()
	if err != nil {
		return fmt.Errorf("opening connection to discord: %w", err)
	}

	// Add commands to discord bot
	log.Info().Msg("adding commands to bot.")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, command := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, cfg.GuildID, command)
		if err != nil {
			return fmt.Errorf("creating command %s: %w", command.Name, err)
		}
		registeredCommands[i] = cmd
	}
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		handler, ok := commandHandlers[i.ApplicationCommandData().Name]
		if ok {
			handler(s, i)
		}
	})

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Info().Msg("Bot is now running.  Press CTRL-C to exit.")
	<-stop

	log.Info().Msg("Shutting down bot. Removing commands.")
	for _, v := range registeredCommands {
		if err := s.ApplicationCommandDelete(s.State.User.ID, cfg.GuildID, v.ID); err != nil {
			log.Err(err).Str("command", v.Name).Msg("failed to delete command")
		}
	}

	return nil
}
