package command

import (
	"fmt"
	"github.com/rs/zerolog/log"

	"github.com/bwmarrin/discordgo"
)

func (c *BotCommands) GetHexaCommand() *discordgo.ApplicationCommand {
	// Workaround since minValue needs to be a pointer to float64
	minValue := 0.0

	var options []*discordgo.ApplicationCommandOption
	options = append(options, &discordgo.ApplicationCommandOption{
		Type:        discordgo.ApplicationCommandOptionInteger,
		Name:        "unused-fragments",
		Description: "Unused Fragments",
		MinValue:    &minValue,
		Required:    true,
	})
	for i := 0; i < c.HexaData.Info.NumOfSkillCore; i++ {
		options = append(options, &discordgo.ApplicationCommandOption{
			Type:        discordgo.ApplicationCommandOptionInteger,
			Name:        fmt.Sprintf("skill-core-%d", i+1),
			Description: fmt.Sprintf("Skill Core %d in top left of Hexa Matrix", i+1),
			MinValue:    &minValue,
			MaxValue:    30,
			Required:    true,
		})
	}
	for i := 0; i < c.HexaData.Info.NumOfBoostCore; i++ {
		options = append(options, &discordgo.ApplicationCommandOption{
			Type:        discordgo.ApplicationCommandOptionInteger,
			Name:        fmt.Sprintf("boost-core-%d", i+1),
			Description: fmt.Sprintf("Boost Core %d in bottom left of Hexa Matrix", i+1),
			MinValue:    &minValue,
			MaxValue:    30,
			Required:    true,
		})
	}
	for i := 0; i < c.HexaData.Info.NumOfMasteryCore; i++ {
		options = append(options, &discordgo.ApplicationCommandOption{
			Type:        discordgo.ApplicationCommandOptionInteger,
			Name:        fmt.Sprintf("mastery-core-%d", i+1),
			Description: fmt.Sprintf("Mastery Core %d in top right Hexa Matrix", i+1),
			MinValue:    &minValue,
			MaxValue:    30,
			Required:    true,
		})
	}
	for i := 0; i < c.HexaData.Info.NumOfCommonCore; i++ {
		options = append(options, &discordgo.ApplicationCommandOption{
			Type:        discordgo.ApplicationCommandOptionInteger,
			Name:        fmt.Sprintf("common-core-%d", i+1),
			Description: fmt.Sprintf("Common Core %d in bottom right Hexa Matrix", i+1),
			MinValue:    &minValue,
			MaxValue:    30,
			Required:    true,
		})
	}

	commands := discordgo.ApplicationCommand{
		Name:        "hexa",
		Description: "calculate total hexa progress",
		Options:     options,
	}

	return &commands
}

func (c *BotCommands) GetHexaCommandHandler() (string, func(s *discordgo.Session, i *discordgo.InteractionCreate), error) {
	hexaCommandHandler := func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options
		optionsMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, option := range options {
			optionsMap[option.Name] = option
		}
		log.Info().Interface("Options", optionsMap).Msg("hexa command options")

		unusedFragments := int(optionsMap["unused-fragments"].IntValue())

		skillCoreFragmentsTotal := 0
		var skillCoreFragments []int
		for i := 0; i < c.HexaData.Info.NumOfSkillCore; i++ {
			skillCoreInput := optionsMap[fmt.Sprintf("skill-core-%d", i+1)].IntValue()
			curTotal := 0
			for level := range skillCoreInput {
				curTotal += c.HexaData.SkillCoreLevelUpChart[level+1]
			}
			skillCoreFragments = append(skillCoreFragments, curTotal)
			skillCoreFragmentsTotal += curTotal

		}

		boostCoreFragmentsTotal := 0
		var boostCoreFragments []int
		for i := 0; i < c.HexaData.Info.NumOfBoostCore; i++ {
			boostCoreInput := optionsMap[fmt.Sprintf("boost-core-%d", i+1)].IntValue()
			curTotal := 0
			for level := range boostCoreInput {
				curTotal += c.HexaData.BoostCoreLevelUpChart[level+1]
			}
			boostCoreFragments = append(boostCoreFragments, curTotal)
			boostCoreFragmentsTotal += curTotal
		}

		masteryCoreFragmentsTotal := 0
		var masteryCoreFragments []int
		for i := 0; i < c.HexaData.Info.NumOfMasteryCore; i++ {
			masteryCoreInput := optionsMap[fmt.Sprintf("mastery-core-%d", i+1)].IntValue()
			curTotal := 0
			for level := range masteryCoreInput {
				curTotal += c.HexaData.MasteryCoreLevelUpChart[level+1]
			}
			masteryCoreFragments = append(masteryCoreFragments, curTotal)
			masteryCoreFragmentsTotal += curTotal
		}

		commonCoreFragmentsTotal := 0
		var commonCoreFragments []int
		for i := 0; i < c.HexaData.Info.NumOfCommonCore; i++ {
			commonCoreInput := optionsMap[fmt.Sprintf("common-core-%d", i+1)].IntValue()
			curTotal := 0
			for level := range commonCoreInput {
				curTotal += c.HexaData.CommonCoreLevelUpChart[level+1]
			}
			commonCoreFragments = append(commonCoreFragments, curTotal)
			commonCoreFragmentsTotal += curTotal
		}

		// Summary
		totalFragmentsFarmed := unusedFragments + skillCoreFragmentsTotal + boostCoreFragmentsTotal + masteryCoreFragmentsTotal + commonCoreFragmentsTotal
		totalFragmentsMaxed := (c.HexaData.Info.NumOfSkillCore * c.HexaData.Info.SkillCoreMaxFragments) +
			(c.HexaData.Info.NumOfBoostCore * c.HexaData.Info.BoostCoreMaxFragments) +
			(c.HexaData.Info.NumOfMasteryCore * c.HexaData.Info.MasteryCoreMaxFragments) +
			(c.HexaData.Info.NumOfCommonCore * c.HexaData.Info.CommonCoreMaxFragments)
		totalFragmentsLeftToMax := totalFragmentsMaxed - totalFragmentsFarmed
		totalFragmentsPercent := (float64(totalFragmentsFarmed) / float64(totalFragmentsMaxed)) * 100
		summary := fmt.Sprintf("**Total Frags:** %d/%d (%.2f%%)\n"+
			"**Frags Needed to Max:** %d\n"+
			"**Unused Frags:** %d\n"+
			"-------------------------------------", totalFragmentsFarmed, totalFragmentsMaxed, totalFragmentsPercent, totalFragmentsLeftToMax, unusedFragments)

		// Breakdown
		skillCoreBreakdown := ""
		for i := range skillCoreFragments {
			skillCoreFragmentsProgress := (float64(skillCoreFragments[i]) / float64(c.HexaData.Info.SkillCoreMaxFragments)) * 100
			skillCoreBreakdown += fmt.Sprintf("**Skill Core %d:**\nLvl: %d | Frags: %d/%d (%.2f%%) | Frags Left: %d\n", i+1, optionsMap[fmt.Sprintf("skill-core-%d", i+1)].IntValue(), skillCoreFragments[i], c.HexaData.Info.SkillCoreMaxFragments, skillCoreFragmentsProgress, c.HexaData.Info.SkillCoreMaxFragments-skillCoreFragments[i])
		}
		skillCoreBreakdown += "-------------------------------------"

		boostCoreBreakdown := ""
		for i := range boostCoreFragments {
			boostCoreFragmentsProgress := (float64(boostCoreFragments[i]) / float64(c.HexaData.Info.BoostCoreMaxFragments)) * 100
			boostCoreBreakdown += fmt.Sprintf("**Boost Core %d:**\nLvl: %d | Frags: %d/%d (%.2f%%) | Frags Left: %d\n", i+1, optionsMap[fmt.Sprintf("boost-core-%d", i+1)].IntValue(), boostCoreFragments[i], c.HexaData.Info.BoostCoreMaxFragments, boostCoreFragmentsProgress, c.HexaData.Info.BoostCoreMaxFragments-boostCoreFragments[i])
		}
		boostCoreBreakdown += "-------------------------------------"

		masteryCoreBreakdown := ""
		for i := range masteryCoreFragments {
			masteryCoreFragmentsProgress := (float64(masteryCoreFragments[i]) / float64(c.HexaData.Info.MasteryCoreMaxFragments)) * 100
			masteryCoreBreakdown += fmt.Sprintf("**Mastery Core %d:**\nLvl: %d | Frags: %d/%d (%.2f%%) | Frags Left: %d\n", i+1, optionsMap[fmt.Sprintf("mastery-core-%d", i+1)].IntValue(), masteryCoreFragments[i], c.HexaData.Info.MasteryCoreMaxFragments, masteryCoreFragmentsProgress, c.HexaData.Info.MasteryCoreMaxFragments-masteryCoreFragments[i])
		}
		masteryCoreBreakdown += "-------------------------------------"

		commonCoreBreakdown := ""
		for i := range commonCoreFragments {
			commonCoreFragmentsProgress := (float64(commonCoreFragments[i]) / float64(c.HexaData.Info.CommonCoreMaxFragments)) * 100
			commonCoreBreakdown += fmt.Sprintf("**Common Core %d:**\nLvl: %d | Frags: %d/%d (%.2f%%) | Frags Left: %d\n", i+1, optionsMap[fmt.Sprintf("common-core-%d", i+1)].IntValue(), commonCoreFragments[i], c.HexaData.Info.CommonCoreMaxFragments, commonCoreFragmentsProgress, c.HexaData.Info.CommonCoreMaxFragments-commonCoreFragments[i])
		}
		commonCoreBreakdown += "---------------------"

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					{
						Title: "Sol Erda Fragment Calculations",
						Fields: []*discordgo.MessageEmbedField{
							{
								Name:  "SUMMARY",
								Value: summary,
							},
							{
								Name:  "SKILL CORE BREAKDOWN",
								Value: skillCoreBreakdown,
							},
							{
								Name:  "BOOST CORE BREAKDOWN",
								Value: boostCoreBreakdown,
							},
							{
								Name:  "MASTERY CORE BREAKDOWN",
								Value: masteryCoreBreakdown,
							},
							{
								Name:  "COMMON CORE BREAKDOWN",
								Value: commonCoreBreakdown,
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

	return "hexa", hexaCommandHandler, nil
}
