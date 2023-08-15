package domain

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestAiGetSkillFromInput(t *testing.T) {
	table := []struct {
		seed int
		want string
	}{
		{0, "普通攻擊"},
		{1, "水球"},
	}

	ai := NewAI(1, "AI", 1000, 100, 100, nil)
	ai.SetSkills(&SkillImpl{NewBasicSkill(ai)}, &SkillImpl{NewWaterBallSkill(ai)}, &SkillImpl{NewCurseSkill(ai)})

	for _, tt := range table {
		ai.setSeed(tt.seed)

		skill := ai.getSkillFromInput()

		if skill.getName() != tt.want {
			t.Errorf("AI should get skill \"基本攻擊\" but got %s", skill.getName())
		}
	}
}

func TestAiGetTargetsFromInput(t *testing.T) {
	ai := NewAI(1, "AI", 1000, 100, 100, nil)
	ai.setSeed(2)

	var candidates []Role
	for i := 1; i <= 3; i++ {
		candidates = append(candidates, NewAI(i, "AI"+strconv.Itoa(i), 1000, 100, 100, nil))
	}

	targets := ai.getTargetsFromInput(candidates, 2)

	assert.Equal(t, 2, len(targets))

	assert.Equal(t, "AI3", targets[0].getName())
	assert.Equal(t, "AI1", targets[1].getName())
	assert.Equal(t, 3, ai.seed)
}
