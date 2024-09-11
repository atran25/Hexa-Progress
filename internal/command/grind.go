package command

import "github.com/bwmarrin/discordgo"

func (c *BotCommands) getGrindCommand() *discordgo.ApplicationCommand {
	var explorerClassChoices []*discordgo.ApplicationCommandOptionChoice
	for _, class := range c.ClassData.Explorer {
		explorerClassChoices = append(explorerClassChoices, &discordgo.ApplicationCommandOptionChoice{
			Name:  class.ClassName,
			Value: class.SlugName,
		})
	}
	grindCommand := discordgo.ApplicationCommand{
		Name:        "grind",
		Description: "Find grinding videos from Korean Maplestory(KMS) for your class",
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
					//{
					//	Name:        "area",
					//	Description: "area name",
					//	Type:        discordgo.ApplicationCommandOptionString,
					//	Required:    true,
					//	Choices:
					//},
				},
			},
		},
	}

	return &grindCommand
}
