package command

import (
	"fmt"

	"github.com/atran25/hexaprogress/internal/hexa"
	"github.com/bwmarrin/discordgo"
)

var (
	// Workaround since minValue needs to b a pointer to float64
	minValue = 0.0

	commands = discordgo.ApplicationCommand{
		Name:        "calc",
		Description: "calculate total hexa progress",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "origin",
				Description: "Origin in hexamatrix",
				MinValue:    &minValue,
				MaxValue:    30,
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "enhance-core-1",
				Description: "Enhance Core 1 in hexa matrix",
				MinValue:    &minValue,
				MaxValue:    30,
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "enhance-core-2",
				Description: "Enhance Core 2 in hexa matrix",
				MinValue:    &minValue,
				MaxValue:    30,
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "enhance-core-3",
				Description: "Enhance Core 3 in hexa matrix",
				MinValue:    &minValue,
				MaxValue:    30,
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "enhance-core-4",
				Description: "Enhance Core 4 in hexa matrix",
				MinValue:    &minValue,
				MaxValue:    30,
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "mastery-core-1",
				Description: "Mastery Core 1 in hexa matrix",
				MinValue:    &minValue,
				MaxValue:    30,
				Required:    true,
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

	originInput := optionsMap["origin"].IntValue()
	originFragments := 0
	for level := range originInput {
		originFragments += hexa.OriginLevelUpChart[level+1]
	}

	enhance1Input := optionsMap["enhance-core-1"].IntValue()
	enhance1Fragments := 0
	for level := range enhance1Input {
		enhance1Fragments += hexa.EnhanceCoreLevelUpChart[level+1]
	}

	enhance2Input := optionsMap["enhance-core-2"].IntValue()
	enhance2Fragments := 0
	for level := range enhance2Input {
		enhance2Fragments += hexa.EnhanceCoreLevelUpChart[level+1]
	}

	enhance3Input := optionsMap["enhance-core-3"].IntValue()
	enhance3Fragments := 0
	for level := range enhance3Input {
		enhance3Fragments += hexa.EnhanceCoreLevelUpChart[level+1]
	}

	enhance4Input := optionsMap["enhance-core-4"].IntValue()
	enhance4Fragments := 0
	for level := range enhance4Input {
		enhance4Fragments += hexa.EnhanceCoreLevelUpChart[level+1]
	}

	mastery1Input := optionsMap["mastery-core-1"].IntValue()
	mastery1Fragments := 0
	for level := range mastery1Input {
		mastery1Fragments += hexa.MasteryCoreLevelUpChart[level+1]
	}

	// Summary
	totalFragments := originFragments + enhance1Fragments + enhance2Fragments + enhance3Fragments + enhance4Fragments + mastery1Fragments
	TotalFragmentsProgress := (float64(totalFragments) / float64(hexa.MaxNumOfFragments)) * 100
	summary := fmt.Sprintf("**Total Fragments:** %d/%d (%.2f%%)\n", totalFragments, hexa.MaxNumOfFragments, TotalFragmentsProgress)

	// Breakdown
	originFragmentsProgress := (float64(originFragments) / float64(hexa.OriginTotal)) * 100
	originBreakdown := fmt.Sprintf("**Origin:**\nLvl: %d | Frags: %d/%d (%.2f%%)\n\n", originInput, originFragments, hexa.OriginTotal, originFragmentsProgress)

	enhance1FragmentsProgress := (float64(enhance1Fragments) / float64(hexa.EnhanceCoreTotal)) * 100
	enhance1Breakdown := fmt.Sprintf("**Enhance Core 1:**\nLvl: %d | Frags: %d/%d (%.2f%%)\n\n", enhance1Input, enhance1Fragments, hexa.EnhanceCoreTotal, enhance1FragmentsProgress)

	enhance2FragmentsProgress := (float64(enhance2Fragments) / float64(hexa.EnhanceCoreTotal)) * 100
	enhance2Breakdown := fmt.Sprintf("**Enhance Core 2:**\nLvl: %d | Frags: %d/%d (%.2f%%)\n\n", enhance2Input, enhance2Fragments, hexa.EnhanceCoreTotal, enhance2FragmentsProgress)

	enhance3FragmentsProgress := (float64(enhance3Fragments) / float64(hexa.EnhanceCoreTotal)) * 100
	enhance3Breakdown := fmt.Sprintf("**Enhance Core 3:**\nLvl: %d | Frags: %d/%d (%.2f%%)\n\n", enhance3Input, enhance3Fragments, hexa.EnhanceCoreTotal, enhance3FragmentsProgress)

	enhance4FragmentsProgress := (float64(enhance4Fragments) / float64(hexa.EnhanceCoreTotal)) * 100
	enhance4Breakdown := fmt.Sprintf("**Enhance Core 4:**\nLvl: %d | Frags: %d/%d (%.2f%%)\n\n", enhance4Input, enhance4Fragments, hexa.EnhanceCoreTotal, enhance4FragmentsProgress)

	mastery1FragmentsProgress := (float64(mastery1Fragments) / float64(hexa.MasteryCoreTotal)) * 100
	mastery1Breakdown := fmt.Sprintf("**Mastery Core 1:**\nLvl: %d | Frags: %d/%d (%.2f%%)\n\n", mastery1Input, mastery1Fragments, hexa.MasteryCoreTotal, mastery1FragmentsProgress)

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
