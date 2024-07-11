package data

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
)

type ClassData struct {
	Explorer []Data `json:"explorer"`
	Cygnus   []Data `json:"cygnus"`
}

type Data struct {
	ClassName  string   `json:"className"`
	SlugName   string   `json:"slugName"`
	KoreanName string   `json:"koreanName"`
	ImageURL   string   `json:"imageURL"`
	Origin     []string `json:"origin"`
	Mastery    []string `json:"mastery"`
	Enhance    []string `json:"enhance"`
	Common     []string `json:"common"`
}

func ReadClassData(path string) (ClassData, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return ClassData{}, fmt.Errorf("reading class.json file: %w", err)
	}

	var classData ClassData
	if err := json.Unmarshal(data, &classData); err != nil {
		return ClassData{}, fmt.Errorf("unmarshalling data into ClassData: %w", err)
	}

	log.Info().Any("classData", classData).Msg("Class data loaded.")

	return classData, nil
}

type BossData struct {
	Difficulty []Difficulty `json:"difficulty"`
	Boss       []Boss       `json:"boss"`
}

type Difficulty struct {
	DifficultyName string `json:"difficultyName"`
	SlugName       string `json:"slugName"`
	KoreanName     string `json:"koreanName"`
}

type Boss struct {
	BossName   string   `json:"bossName"`
	SlugName   string   `json:"slugName"`
	KoreanName string   `json:"koreanName"`
	Difficulty []string `json:"difficulty"`
	ImageURL   string   `json:"imageURL"`
}

func ReadBossData(path string) (BossData, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return BossData{}, fmt.Errorf("reading boss.json file: %w", err)
	}

	var bossData BossData
	if err := json.Unmarshal(data, &bossData); err != nil {
		return BossData{}, fmt.Errorf("unmarshalling data into []BossData: %w", err)
	}

	log.Info().Any("bossData", bossData).Msg("Boss data loaded.")

	return bossData, nil
}
