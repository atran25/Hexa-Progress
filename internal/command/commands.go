package command

import "github.com/atran25/hexaprogress/data"

type CommandHandler struct {
	ClassDataMap  map[string]data.Data
	BossDataMap   map[string]data.Boss
	DifficultyMap map[string]data.Difficulty
	ClassData     data.ClassData
	BossData      data.BossData
}

func NewCommandHandler(classData data.ClassData, bossData data.BossData) *CommandHandler {
	classDataMap := make(map[string]data.Data)
	for _, class := range classData.Explorer {
		classDataMap[class.SlugName] = class
	}
	for _, class := range classData.Cygnus {
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

	return &CommandHandler{
		ClassDataMap:  classDataMap,
		BossDataMap:   bossDataMap,
		DifficultyMap: difficultyMap,
		ClassData:     classData,
		BossData:      bossData,
	}
}
