package command

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/gosimple/slug"
	"github.com/rs/zerolog/log"
)

func (c *CommandHandler) GetBossCommand() *discordgo.ApplicationCommand {
	var explorerClassChoices []*discordgo.ApplicationCommandOptionChoice
	for _, class := range c.ClassData.Explorer {
		explorerClassChoices = append(explorerClassChoices, &discordgo.ApplicationCommandOptionChoice{
			Name:  class.ClassName,
			Value: class.SlugName,
		})
	}

	var cygnusClassChoices []*discordgo.ApplicationCommandOptionChoice
	for _, class := range c.ClassData.Cygnus {
		cygnusClassChoices = append(cygnusClassChoices, &discordgo.ApplicationCommandOptionChoice{
			Name:  class.ClassName,
			Value: class.SlugName,
		})
	}

	var bossChoices []*discordgo.ApplicationCommandOptionChoice
	for _, boss := range c.BossData.Boss {
		bossChoices = append(bossChoices, &discordgo.ApplicationCommandOptionChoice{
			Name:  boss.BossName,
			Value: boss.SlugName,
		})
	}

	bossCommand := discordgo.ApplicationCommand{
		Name:        "boss",
		Description: "Find bossing videos from Korean Maplestory(KMS) for your class",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        "explorer",
				Description: "explorer branch",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "class",
						Description: "explorer class",
						Type:        discordgo.ApplicationCommandOptionString,
						Required:    true,
						Choices:     explorerClassChoices,
					},
					{
						Name:        "boss",
						Description: "boss name",
						Type:        discordgo.ApplicationCommandOptionString,
						Required:    true,
						Choices:     bossChoices,
					},
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        "cygnus",
				Description: "cygnus branch",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "class",
						Description: "cygnus class",
						Type:        discordgo.ApplicationCommandOptionString,
						Required:    true,
						Choices:     cygnusClassChoices,
					},
					{
						Name:        "boss",
						Description: "boss name",
						Type:        discordgo.ApplicationCommandOptionString,
						Required:    true,
						Choices:     bossChoices,
					},
				},
			},
		},
	}

	return &bossCommand
}

func (c *CommandHandler) GetBossCommandHandler() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	bossCommandHandler := map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"boss": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			options := i.ApplicationCommandData().Options[0].Options
			log.Info().Interface("Options", options).Msg("options")

			/*
			*
			 */
			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			for _, option := range options {
				optionMap[option.Name] = option
			}

			class := c.ClassDataMap[optionMap["class"].StringValue()]
			boss := c.BossDataMap[optionMap["boss"].StringValue()]

			searchChoices := fmt.Sprintf("**Class Name:** %s\n**Boss Name:** %s\n**%s** has **%d** difficulties", class.ClassName, boss.BossName, boss.BossName, len(boss.Difficulty))
			searchResults := ""
			for _, currDiff := range boss.Difficulty {
				youtubeURL := fmt.Sprintf("https://www.youtube.com/results?search_query=%s+%s+%s", class.KoreanName, c.DifficultyMap[slug.Make(currDiff)].KoreanName, boss.KoreanName)
				youtubeURL = strings.ReplaceAll(youtubeURL, " ", "+")
				searchResults += fmt.Sprintf("\n**%s %s %s:**\n %s", class.ClassName, currDiff, boss.BossName, youtubeURL)
			}
			youtubeURL := fmt.Sprintf("https://www.youtube.com/results?search_query=%s+%s", class.KoreanName, boss.KoreanName)
			youtubeURL = strings.ReplaceAll(youtubeURL, " ", "+")
			searchResults += fmt.Sprintf("\n**All difficulties:**\n %s", youtubeURL)
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						{
							Title: fmt.Sprintf("Korean Maplestory Bossing Videos (%s)", class.ClassName),
							Thumbnail: &discordgo.MessageEmbedThumbnail{
								URL: class.ImageURL,
							},
							Image: &discordgo.MessageEmbedImage{
								URL: boss.ImageURL,
							},
							Fields: []*discordgo.MessageEmbedField{
								{
									Name:  "Search Formula",
									Value: "Class Name + Boss Difficulty + Boss Name",
								},
								{
									Value: searchChoices,
								},
								{
									Value: searchResults,
								},
							},
							Footer: &discordgo.MessageEmbedFooter{
								Text: "Free bot, if you paid for this you got scammed \nSource code: https://github.com/atran25/Hexa-Progress",
							},
						},
					},
				},
			})
		},
	}

	return bossCommandHandler
}
