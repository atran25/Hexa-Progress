package main

import (
	"github.com/caarlos0/env/v11"
)

type config struct {
	DiscordToken string `env:"HEXAPROGRESS_DISCORD_TOKEN"`
	GuildID      string `env:"HEXAPROGRESS_GUILD_ID"`
}

func GetConfig() (config, error) {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		return config{}, err
	}
	return cfg, nil
}
