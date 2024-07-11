package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadClassData(t *testing.T) {
	classData, err := ReadClassData("class.json")
	assert.NoErrorf(t, err, "ReadClassData() error = %v", err)
	assert.NotEmpty(t, classData)
	assert.Len(t, classData.Explorer, 2)
}

func TestReadBossData(t *testing.T) {
	bossData, err := ReadBossData("boss.json")
	assert.NoErrorf(t, err, "ReadBossData() error = %v", err)
	assert.NotEmpty(t, bossData)
	assert.Len(t, bossData.Difficulty, 5)
	assert.Len(t, bossData.Boss, 12)
}
