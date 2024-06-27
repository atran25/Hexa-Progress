package main

import (
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

var (
	// Workaround since minValue needs to b a pointer to float64
	minValue = float64(0)

	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "calc",
			Description: "calculate total hexa progress",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "origin",
					Description: "Origin in hexamatrix",
					Required:    true,
					MinValue:    &minValue,
					MaxValue:    30,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "enhance-core-1",
					Description: "Enhance Core 1 in hexa matrix",
					Required:    true,
					MinValue:    &minValue,
					MaxValue:    30,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "enhance-core-2",
					Description: "Enhance Core 2 in hexa matrix",
					Required:    true,
					MinValue:    &minValue,
					MaxValue:    30,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "enhance-core-3",
					Description: "Enhance Core 3 in hexa matrix",
					Required:    true,
					MinValue:    &minValue,
					MaxValue:    30,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "enhance-core-4",
					Description: "Enhance Core 4 in hexa matrix",
					Required:    true,
					MinValue:    &minValue,
					MaxValue:    30,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "mastery-core-1",
					Description: "Mastery Core 1 in hexa matrix",
					Required:    true,
					MinValue:    &minValue,
					MaxValue:    30,
				},
			},
		},
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"calc": CalcResponse,
	}
)

func CalcResponse(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	optionsMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, option := range options {
		optionsMap[option.Name] = option
	}

	originInput, _ := strconv.Atoi(optionsMap["origin"].StringValue())
	originFragments := 0
	for level := range originInput {
		originFragments += originLevelUpChart[level+1]
	}

	enhance1Input, _ := strconv.Atoi(optionsMap["enhance-core-1"].StringValue())
	enhance1Fragments := 0
	for level := range enhance1Input {
		enhance1Fragments += enhanceCoreLevelUpChart[level+1]
	}

	enhance2Input, _ := strconv.Atoi(optionsMap["enhance-core-2"].StringValue())
	enhance2Fragments := 0
	for level := range enhance2Input {
		enhance2Fragments += enhanceCoreLevelUpChart[level+1]
	}

	enhance3Input, _ := strconv.Atoi(optionsMap["enhance-core-3"].StringValue())
	enhance3Fragments := 0
	for level := range enhance3Input {
		enhance3Fragments += enhanceCoreLevelUpChart[level+1]
	}

	enhance4Input, _ := strconv.Atoi(optionsMap["enhance-core-4"].StringValue())
	enhance4Fragments := 0
	for level := range enhance4Input {
		enhance4Fragments += enhanceCoreLevelUpChart[level+1]
	}

	mastery1Input, _ := strconv.Atoi(optionsMap["mastery-core-1"].StringValue())
	mastery1Fragments := 0
	for level := range mastery1Input {
		mastery1Fragments += masteryCoreLevelUpChart[level+1]
	}

	// Summary
	totalFragments := originFragments + enhance1Fragments + enhance2Fragments + enhance3Fragments + enhance4Fragments + mastery1Fragments
	TotalFragmentsProgress := (float64(totalFragments) / float64(MaxNumOfFragments)) * 100
	summary := fmt.Sprintf("**Total Fragments:** %d/%d (%.2f%%)\n", totalFragments, MaxNumOfFragments, TotalFragmentsProgress)

	// Breakdown
	originFragmentsProgress := (float64(originFragments) / float64(originTotal)) * 100
	originBreakdown := fmt.Sprintf("**Origin:**\nLvl: %d | Frags: %d/%d (%.2f%%)\n\n", originInput, originFragments, originTotal, originFragmentsProgress)

	enhance1FragmentsProgress := (float64(enhance1Fragments) / float64(enhanceCoreTotal)) * 100
	enhance1Breakdown := fmt.Sprintf("**Enhance Core 1:**\nLvl: %d | Frags: %d/%d (%.2f%%)\n\n", enhance1Input, enhance1Fragments, enhanceCoreTotal, enhance1FragmentsProgress)

	enhance2FragmentsProgress := (float64(enhance2Fragments) / float64(enhanceCoreTotal)) * 100
	enhance2Breakdown := fmt.Sprintf("**Enhance Core 2:**\nLvl: %d | Frags: %d/%d (%.2f%%)\n\n", enhance2Input, enhance2Fragments, enhanceCoreTotal, enhance2FragmentsProgress)

	enhance3FragmentsProgress := (float64(enhance3Fragments) / float64(enhanceCoreTotal)) * 100
	enhance3Breakdown := fmt.Sprintf("**Enhance Core 3:**\nLvl: %d | Frags: %d/%d (%.2f%%)\n\n", enhance3Input, enhance3Fragments, enhanceCoreTotal, enhance3FragmentsProgress)

	enhance4FragmentsProgress := (float64(enhance4Fragments) / float64(enhanceCoreTotal)) * 100
	enhance4Breakdown := fmt.Sprintf("**Enhance Core 4:**\nLvl: %d | Frags: %d/%d (%.2f%%)\n\n", enhance4Input, enhance4Fragments, enhanceCoreTotal, enhance4FragmentsProgress)

	mastery1FragmentsProgress := (float64(mastery1Fragments) / float64(masteryCoreTotal)) * 100
	mastery1Breakdown := fmt.Sprintf("**Mastery Core 1:**\nLvl: %d | Frags: %d/%d (%.2f%%)\n\n", mastery1Input, mastery1Fragments, masteryCoreTotal, mastery1FragmentsProgress)

	breakdown := originBreakdown + enhance1Breakdown + enhance2Breakdown + enhance3Breakdown + enhance4Breakdown + mastery1Breakdown

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title: "Sol Erda Fragment Calculations",
					Fields: []*discordgo.MessageEmbedField{
						{
							Name:  "Summary",
							Value: summary,
						},
						{
							Name:  "Breakdown",
							Value: breakdown,
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
