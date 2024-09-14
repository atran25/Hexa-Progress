package command

import (
	"fmt"

	"github.com/atran25/hexaprogress/data"
	"github.com/bwmarrin/discordgo"
)

type BotCommands struct {
	ClassDataMap  map[string]data.Data
	BossDataMap   map[string]data.Boss
	DifficultyMap map[string]data.Difficulty
	ClassData     data.ClassData
	BossData      data.BossData
	HexaData      data.HexaData
}

func NewBotCommands(classData data.ClassData, bossData data.BossData, hexaData data.HexaData) *BotCommands {
	classDataMap := make(map[string]data.Data)
	for _, class := range classData.Explorer {
		classDataMap[class.SlugName] = class
	}
	for _, class := range classData.Cygnus {
		classDataMap[class.SlugName] = class
	}
	for _, class := range classData.Hero {
		classDataMap[class.SlugName] = class
	}
	for _, class := range classData.Resistance {
		classDataMap[class.SlugName] = class
	}
	for _, class := range classData.Nova {
		classDataMap[class.SlugName] = class
	}
	for _, class := range classData.Flora {
		classDataMap[class.SlugName] = class
	}
	for _, class := range classData.Anima {
		classDataMap[class.SlugName] = class
	}
	for _, class := range classData.Other {
		classDataMap[class.SlugName] = class
	}
	bossDataMap := make(map[string]data.Boss)
	for _, boss := range bossData.Boss {
		bossDataMap[boss.SlugName] = boss
	}
	difficultyMap := make(map[string]data.Difficulty)
	for _, difficulty := range bossData.Difficulty {
		difficultyMap[difficulty.SlugName] = difficulty
	}

	return &BotCommands{
		ClassDataMap:  classDataMap,
		BossDataMap:   bossDataMap,
		DifficultyMap: difficultyMap,
		ClassData:     classData,
		BossData:      bossData,
		HexaData:      hexaData,
	}
}

func (c *BotCommands) GetAllCommands() []*discordgo.ApplicationCommand {
	return []*discordgo.ApplicationCommand{
		c.GetBossCommand(),
		c.GetHexaCommand(),
	}
}

func (c *BotCommands) GetAllCommandsHandler() (map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate), error) {
	bossCommandName, bossCommandHandler, err := c.GetBossCommandHandler()
	if err != nil {
		return nil, fmt.Errorf("getting boss command: %w", err)
	}

	hexaCommandName, hexaCommandHandler, err := c.GetHexaCommandHandler()
	if err != nil {
		return nil, fmt.Errorf("getting hexa command: %w", err)
	}

	return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		bossCommandName: bossCommandHandler,
		hexaCommandName: hexaCommandHandler,
	}, nil

}
