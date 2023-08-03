package domain

import (
	"C3M3H1/domain/enum"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCharacterAttackTwoMonstersThenMonstersDieScenario(t *testing.T) {
	m := NewMap(10, NewGame())
	character := NewCharacter(m)
	character.setDirection(enum.Down)
	m.putRoleAt(character, 2, 3)
	m1 := NewMonster(m)
	m2 := NewMonster(m)
	m.putRoleAt(m1, 5, 3)
	m.putRoleAt(m2, 6, 3)

	character.attack()

	assert.Equal(t, 1, len(m.roles))

	findMonster := false

	for i := 0; i < len(m.objects); i++ {
		for j := 0; j < len(m.objects); j++ {
			if m.objects[i][j] == m1 || m.objects[i][j] == m2 {
				findMonster = true
			}
		}
	}

	assert.False(t, findMonster)
}

func TestCharacterAttackTwoMonstersWhenOneOfMonsterStateInvincibleThenMonstersDieScenario(t *testing.T) {
	m := NewMap(10, NewGame())
	character := NewCharacter(m)
	character.setDirection(enum.Down)
	m.putRoleAt(character, 2, 3)
	m1 := NewMonster(m)
	m2 := NewMonster(m)
	m2.applyState(NewInvincible(m2))
	m.putRoleAt(m1, 5, 3)
	m.putRoleAt(m2, 6, 3)

	character.attack()

	assert.Equal(t, 2, len(m.roles))

	findM2 := false

	for i := 0; i < len(m.objects); i++ {
		for j := 0; j < len(m.objects); j++ {
			if m.objects[i][j] == m2 {
				findM2 = true
			}
		}
	}

	assert.True(t, findM2)
}
