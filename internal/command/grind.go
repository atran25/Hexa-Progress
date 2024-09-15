package command

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"strings"
)

func (c *BotCommands) GetGrindCommand() *discordgo.ApplicationCommand {
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

	var heroClassChoices []*discordgo.ApplicationCommandOptionChoice
	for _, class := range c.ClassData.Hero {
		heroClassChoices = append(heroClassChoices, &discordgo.ApplicationCommandOptionChoice{
			Name:  class.ClassName,
			Value: class.SlugName,
		})
	}

	var resistanceClassChoices []*discordgo.ApplicationCommandOptionChoice
	for _, class := range c.ClassData.Resistance {
		resistanceClassChoices = append(resistanceClassChoices, &discordgo.ApplicationCommandOptionChoice{
			Name:  class.ClassName,
			Value: class.SlugName,
		})
	}

	var novaClassChoices []*discordgo.ApplicationCommandOptionChoice
	for _, class := range c.ClassData.Nova {
		novaClassChoices = append(novaClassChoices, &discordgo.ApplicationCommandOptionChoice{
			Name:  class.ClassName,
			Value: class.SlugName,
		})
	}

	var floraClassChoices []*discordgo.ApplicationCommandOptionChoice
	for _, class := range c.ClassData.Flora {
		floraClassChoices = append(floraClassChoices, &discordgo.ApplicationCommandOptionChoice{
			Name:  class.ClassName,
			Value: class.SlugName,
		})
	}

	var animaClassChoices []*discordgo.ApplicationCommandOptionChoice
	for _, class := range c.ClassData.Anima {
		animaClassChoices = append(animaClassChoices, &discordgo.ApplicationCommandOptionChoice{
			Name:  class.ClassName,
			Value: class.SlugName,
		})
	}

	var otherClassChoices []*discordgo.ApplicationCommandOptionChoice
	for _, class := range c.ClassData.Other {
		otherClassChoices = append(otherClassChoices, &discordgo.ApplicationCommandOptionChoice{
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

	var arcaneAreaChoices []*discordgo.ApplicationCommandOptionChoice
	for _, area := range c.AreaData.ArcaneRiver {
		arcaneAreaChoices = append(arcaneAreaChoices, &discordgo.ApplicationCommandOptionChoice{
			Name:  area.Name,
			Value: area.SlugName,
		})
	}

	var grandisAreaChoices []*discordgo.ApplicationCommandOptionChoice
	for _, area := range c.AreaData.Grandis {
		grandisAreaChoices = append(grandisAreaChoices, &discordgo.ApplicationCommandOptionChoice{
			Name:  area.Name,
			Value: area.SlugName,
		})
	}

	grindCommand := discordgo.ApplicationCommand{
		Name:        "grind",
		Description: "Find grinding videos from Korean Maplestory(KMS) for your class",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionSubCommandGroup,
				Name:        "arcane",
				Description: "arcane river grinding",
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
								Name:        "area",
								Description: "arcane river area",
								Type:        discordgo.ApplicationCommandOptionString,
								Required:    true,
								Choices:     arcaneAreaChoices,
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
								Name:        "area",
								Description: "arcane river area",
								Type:        discordgo.ApplicationCommandOptionString,
								Required:    true,
								Choices:     arcaneAreaChoices,
							},
						},
					},
					{
						Type:        discordgo.ApplicationCommandOptionSubCommand,
						Name:        "hero",
						Description: "hero branch",
						Options: []*discordgo.ApplicationCommandOption{
							{
								Name:        "class",
								Description: "hero class",
								Type:        discordgo.ApplicationCommandOptionString,
								Required:    true,
								Choices:     heroClassChoices,
							},
							{
								Name:        "area",
								Description: "arcane river area",
								Type:        discordgo.ApplicationCommandOptionString,
								Required:    true,
								Choices:     arcaneAreaChoices,
							},
						},
					},
					{
						Type:        discordgo.ApplicationCommandOptionSubCommand,
						Name:        "resistance",
						Description: "resistance branch",
						Options: []*discordgo.ApplicationCommandOption{
							{
								Name:        "class",
								Description: "resistance class",
								Type:        discordgo.ApplicationCommandOptionString,
								Required:    true,
								Choices:     resistanceClassChoices,
							},
							{
								Name:        "area",
								Description: "arcane river area",
								Type:        discordgo.ApplicationCommandOptionString,
								Required:    true,
								Choices:     arcaneAreaChoices,
							},
						},
					},
					{
						Type:        discordgo.ApplicationCommandOptionSubCommand,
						Name:        "nova",
						Description: "nova branch",
						Options: []*discordgo.ApplicationCommandOption{
							{
								Name:        "class",
								Description: "nova class",
								Type:        discordgo.ApplicationCommandOptionString,
								Required:    true,
								Choices:     novaClassChoices,
							},
							{
								Name:        "area",
								Description: "arcane river area",
								Type:        discordgo.ApplicationCommandOptionString,
								Required:    true,
								Choices:     arcaneAreaChoices,
							},
						},
					},
					{
						Type:        discordgo.ApplicationCommandOptionSubCommand,
						Name:        "flora",
						Description: "flora branch",
						Options: []*discordgo.ApplicationCommandOption{
							{
								Name:        "class",
								Description: "flora class",
								Type:        discordgo.ApplicationCommandOptionString,
								Required:    true,
								Choices:     floraClassChoices,
							},
							{
								Name:        "area",
								Description: "arcane river area",
								Type:        discordgo.ApplicationCommandOptionString,
								Required:    true,
								Choices:     arcaneAreaChoices,
							},
						},
					},
					{
						Type:        discordgo.ApplicationCommandOptionSubCommand,
						Name:        "anima",
						Description: "anima branch",
						Options: []*discordgo.ApplicationCommandOption{
							{
								Name:        "class",
								Description: "anima class",
								Type:        discordgo.ApplicationCommandOptionString,
								Required:    true,
								Choices:     animaClassChoices,
							},
							{
								Name:        "area",
								Description: "arcane river area",
								Type:        discordgo.ApplicationCommandOptionString,
								Required:    true,
								Choices:     arcaneAreaChoices,
							},
						},
					},
					{
						Type:        discordgo.ApplicationCommandOptionSubCommand,
						Name:        "other",
						Description: "other branch",
						Options: []*discordgo.ApplicationCommandOption{
							{
								Name:        "class",
								Description: "other class",
								Type:        discordgo.ApplicationCommandOptionString,
								Required:    true,
								Choices:     otherClassChoices,
							},
							{
								Name:        "area",
								Description: "arcane river area",
								Type:        discordgo.ApplicationCommandOptionString,
								Required:    true,
								Choices:     arcaneAreaChoices,
							},
						},
					},
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionSubCommandGroup,
				Name:        "grandis",
				Description: "grandis grinding",
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
								Name:        "area",
								Description: "grandis area",
								Type:        discordgo.ApplicationCommandOptionString,
								Required:    true,
								Choices:     grandisAreaChoices,
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
								Name:        "area",
								Description: "grandis area",
								Type:        discordgo.ApplicationCommandOptionString,
								Required:    true,
								Choices:     grandisAreaChoices,
							},
						},
					},
					{
						Type:        discordgo.ApplicationCommandOptionSubCommand,
						Name:        "hero",
						Description: "hero branch",
						Options: []*discordgo.ApplicationCommandOption{
							{
								Name:        "class",
								Description: "hero class",
								Type:        discordgo.ApplicationCommandOptionString,
								Required:    true,
								Choices:     heroClassChoices,
							},
							{
								Name:        "area",
								Description: "grandis area",
								Type:        discordgo.ApplicationCommandOptionString,
								Required:    true,
								Choices:     grandisAreaChoices,
							},
						},
					},
					{
						Type:        discordgo.ApplicationCommandOptionSubCommand,
						Name:        "resistance",
						Description: "resistance branch",
						Options: []*discordgo.ApplicationCommandOption{
							{
								Name:        "class",
								Description: "resistance class",
								Type:        discordgo.ApplicationCommandOptionString,
								Required:    true,
								Choices:     resistanceClassChoices,
							},
							{
								Name:        "area",
								Description: "grandis area",
								Type:        discordgo.ApplicationCommandOptionString,
								Required:    true,
								Choices:     grandisAreaChoices,
							},
						},
					},
					{
						Type:        discordgo.ApplicationCommandOptionSubCommand,
						Name:        "nova",
						Description: "nova branch",
						Options: []*discordgo.ApplicationCommandOption{
							{
								Name:        "class",
								Description: "nova class",
								Type:        discordgo.ApplicationCommandOptionString,
								Required:    true,
								Choices:     novaClassChoices,
							},
							{
								Name:        "area",
								Description: "grandis area",
								Type:        discordgo.ApplicationCommandOptionString,
								Required:    true,
								Choices:     grandisAreaChoices,
							},
						},
					},
					{
						Type:        discordgo.ApplicationCommandOptionSubCommand,
						Name:        "flora",
						Description: "flora branch",
						Options: []*discordgo.ApplicationCommandOption{
							{
								Name:        "class",
								Description: "flora class",
								Type:        discordgo.ApplicationCommandOptionString,
								Required:    true,
								Choices:     floraClassChoices,
							},
							{
								Name:        "area",
								Description: "grandis area",
								Type:        discordgo.ApplicationCommandOptionString,
								Required:    true,
								Choices:     grandisAreaChoices,
							},
						},
					},
					{
						Type:        discordgo.ApplicationCommandOptionSubCommand,
						Name:        "anima",
						Description: "anima branch",
						Options: []*discordgo.ApplicationCommandOption{
							{
								Name:        "class",
								Description: "anima class",
								Type:        discordgo.ApplicationCommandOptionString,
								Required:    true,
								Choices:     animaClassChoices,
							},
							{
								Name:        "area",
								Description: "grandis area",
								Type:        discordgo.ApplicationCommandOptionString,
								Required:    true,
								Choices:     grandisAreaChoices,
							},
						},
					},
					{
						Type:        discordgo.ApplicationCommandOptionSubCommand,
						Name:        "other",
						Description: "other branch",
						Options: []*discordgo.ApplicationCommandOption{
							{
								Name:        "class",
								Description: "other class",
								Type:        discordgo.ApplicationCommandOptionString,
								Required:    true,
								Choices:     otherClassChoices,
							},
							{
								Name:        "area",
								Description: "grandis area",
								Type:        discordgo.ApplicationCommandOptionString,
								Required:    true,
								Choices:     grandisAreaChoices,
							},
						},
					},
				},
			},
		},
	}

	return &grindCommand
}

func (c *BotCommands) GetGrindCommandHandler() (string, func(s *discordgo.Session, i *discordgo.InteractionCreate), error) {
	grindCommandHandler := func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// The options are double nested because we are using a double sub command
		options := i.ApplicationCommandData().Options[0].Options[0].Options
		log.Info().Interface("ApplicationCommandData", i.ApplicationCommandData()).Msg("ApplicationCommandData")
		log.Info().Interface("Options", options).Msg("options")

		optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, option := range options {
			optionMap[option.Name] = option
		}
		log.Info().Interface("OptionMap", optionMap).Msg("optionMap")

		class := c.ClassDataMap[optionMap["class"].StringValue()]
		area := c.AreaDataMap[optionMap["area"].StringValue()]

		searchChoices := fmt.Sprintf("**Class Name:** %s\n**Area:** %s\n**%s** has **%d** sub areas", class.ClassName, area.Name, area.Name, len(area.SubArea))

		searchResults := ""
		for _, subArea := range area.SubArea {
			youtubeURL := fmt.Sprintf("https://www.youtube.com/results?search_query=%s+%s+%s", class.KoreanName, area.KoreanName, subArea.KoreanName)
			youtubeURL = strings.ReplaceAll(youtubeURL, " ", "+")
			searchResults += fmt.Sprintf("\n**%s %s %s:**\n %s", class.ClassName, area.Name, subArea.Name, youtubeURL)
		}
		youtubeURL := fmt.Sprintf("https://www.youtube.com/results?search_query=%s+%s", class.KoreanName, area.KoreanName)
		youtubeURL = strings.ReplaceAll(youtubeURL, " ", "+")
		searchResults += fmt.Sprintf("\n**All sub areas:**\n %s", youtubeURL)

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					{
						Title: fmt.Sprintf("Korean Maplestory Grinding Videos (%s)", class.ClassName),
						Thumbnail: &discordgo.MessageEmbedThumbnail{
							URL: class.ImageURL,
						},
						Image: &discordgo.MessageEmbedImage{
							URL: area.ImageURL,
						},
						Fields: []*discordgo.MessageEmbedField{
							{
								Name:  "Search Formula",
								Value: "Class Name + Area Name",
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
	}
	return "grind", grindCommandHandler, nil
}
