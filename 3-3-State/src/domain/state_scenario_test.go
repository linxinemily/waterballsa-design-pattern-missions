package domain

import (
	"C3M3H1/domain/enum"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInvincibleStateScenario(t *testing.T) {
	m := NewMap(10, NewGame())
	character := NewCharacter(m)
	character.applyState(NewInvincible(character))

	assert.True(t, character.isFullHP())

	for i := 0; i < 2; i++ { // 經過 2 回合
		assert.False(t, character.checkStateIfExpired())

		assert.Equal(t, 1, character.beforeTakeTurn())

		character.attacked() // 期間遭受攻擊
		assert.True(t, character.isFullHP())

		character.roundEnd()
	}

	assert.True(t, character.checkStateIfExpired())
	assert.Equal(t, NewNormal(character).String(), character.state.String())
}

func TestPoisonedStateScenario(t *testing.T) {
	t.Run("主角", func(t *testing.T) {
		m := NewMap(10, NewGame())
		character := NewCharacter(m)
		character.applyState(NewPoisoned(character))

		assertHP := character.fullHP
		for i := 0; i < 3; i++ { // 經過 3 回合
			assert.False(t, character.checkStateIfExpired())

			assert.Equal(t, 1, character.beforeTakeTurn())

			assertHP -= 15
			assert.Equal(t, assertHP, character.hp)

			character.roundEnd()
		}

		assert.True(t, character.checkStateIfExpired())
		assert.Equal(t, NewNormal(character).String(), character.state.String())

	})

	t.Run("怪物", func(t *testing.T) {
		m := NewMap(10, NewGame())
		monster := NewMonster(m)
		m.putRoleAt(monster, 5, 3)
		monster.applyState(NewPoisoned(monster))
		assert.False(t, monster.checkStateIfExpired())
		assert.Equal(t, 1, monster.beforeTakeTurn())
		monster.roundEnd()

		res, _ := monster.m.getObjectAt(5, 3)
		assert.True(t, res == nil)
		assert.True(t, len(m.roles) == 0)
	})

}

func TestAcceleratedStateScenario(t *testing.T) {
	m := NewMap(10, NewGame())
	character := NewCharacter(m)

	t.Run("期間未遭受攻擊", func(t *testing.T) {
		character.applyState(NewAccelerated(character))

		for i := 0; i < 3; i++ { // 經過 3 回合
			assert.False(t, character.checkStateIfExpired())

			assert.Equal(t, 2, character.beforeTakeTurn())

			character.roundEnd()
		}

		assert.Equal(t, character.state.String(), NewAccelerated(character).String()) // 仍是加速狀態
		assert.True(t, character.checkStateIfExpired())                               // 加速狀態已過期
		assert.Equal(t, NewNormal(character).String(), character.state.String())
	})

	t.Run("在期間遭受攻擊，回復成一般狀態", func(t *testing.T) {
		character.applyState(NewAccelerated(character))

		for i := 0; i < 2; i++ {
			assert.False(t, character.checkStateIfExpired())

			assert.Equal(t, 2, character.beforeTakeTurn())

			if i == 1 { // 執行第 2 次時遭受攻擊
				character.attacked()
			}

			character.roundEnd()
		}

		assert.False(t, character.checkStateIfExpired()) // 已經變回正常狀態
		assert.Equal(t, NewNormal(character).String(), character.state.String())
	})

}

func TestHealingStateScenario(t *testing.T) {
	m := NewMap(10, NewGame())
	character := NewCharacter(m)

	t.Run("原本只有 100 HP，加到 100 + 30*5 = 250", func(t *testing.T) {
		character.applyState(NewHealing(character))
		character.hp = 100
		assertHP := character.hp
		for i := 0; i < 5; i++ { // 經過 5 回合
			assert.False(t, character.checkStateIfExpired())

			assert.Equal(t, 1, character.beforeTakeTurn())

			assertHP += 30
			assert.Equal(t, assertHP, character.hp)

			character.roundEnd()
		}

		assert.True(t, character.checkStateIfExpired())
		assert.Equal(t, NewNormal(character).String(), character.state.String())
	})

	t.Run("原本有 200 HP，加第 4 次的時候就滿血了，回復成正常狀態", func(t *testing.T) {
		character.applyState(NewHealing(character))
		character.hp = 200
		assertHP := character.hp
		for i := 0; i < 4; i++ {
			assert.False(t, character.checkStateIfExpired())

			assert.Equal(t, 1, character.beforeTakeTurn())

			if i == 3 { // 加第 4 次的時候就滿血了
				assertHP = character.fullHP
			} else {
				assertHP += 30
			}

			assert.Equal(t, assertHP, character.hp)

			character.roundEnd()
		}

		assert.False(t, character.checkStateIfExpired()) //已回復成正常狀態
		assert.Equal(t, NewNormal(character).String(), character.state.String())
	})

}

func TestOrderlessStateScenario(t *testing.T) {
	m := NewMap(10, NewGame())
	character := NewCharacter(m)
	character.applyState(NewOrderless(character))
	character.hp = 100
	for i := 0; i < 3; i++ { // 經過 3 回合
		assert.False(t, character.checkStateIfExpired())

		assert.Equal(t, 1, character.beforeTakeTurn())

		assert.Equal(t, 2, len(character.state.getValidDirections()))

		character.roundEnd()
	}

	assert.True(t, character.checkStateIfExpired())
	assert.Equal(t, NewNormal(character).String(), character.state.String())
	assert.Equal(t, 4, len(character.state.getValidDirections()))
}

func TestStockpileStateScenario(t *testing.T) {
	m := NewMap(10, NewGame())
	character := NewCharacter(m)

	t.Run("兩回合後進入爆發狀態", func(t *testing.T) {
		character.applyState(NewStockpile(character))
		for i := 0; i < 2; i++ { // 經過 2 回合
			assert.False(t, character.checkStateIfExpired())

			assert.Equal(t, 1, character.beforeTakeTurn())

			character.roundEnd()
		}

		assert.True(t, character.checkStateIfExpired())
		assert.Equal(t, NewErupting(character).String(), character.state.String())
	})

	t.Run("第一回合遭受攻擊，回復成一般狀態", func(t *testing.T) {
		character.applyState(NewHealing(character))
		character.applyState(NewStockpile(character))
		for i := 0; i < 2; i++ {
			assert.False(t, character.checkStateIfExpired())

			assert.Equal(t, 1, character.beforeTakeTurn())

			if i == 0 {
				character.attacked() //第一回合遭受攻擊
			}

			character.roundEnd()
		}

		assert.Equal(t, character.fullHP, character.hp)  //第一回合攻擊無效(被攻擊完立即恢復成一般狀態)
		assert.False(t, character.checkStateIfExpired()) //一般狀態不會過期
		assert.Equal(t, NewNormal(character).String(), character.state.String())
	})
}

func TestEruptingStateScenario(t *testing.T) {
	t.Run("主角", func(t *testing.T) {
		m := NewMap(10, NewGame())
		character := NewCharacter(m)
		character.setDirection(enum.Down)
		character.applyState(NewErupting(character))
		m.putRoleAt(character, 2, 3)
		m1 := NewMonster(m)
		m2 := NewMonster(m)
		// 在 character 原本的攻擊範圍內
		m.putRoleAt(m1, 5, 3)
		m.putRoleAt(m2, 6, 3)
		// 不在 character 原本的攻擊範圍內
		m3 := NewMonster(m)
		m.putRoleAt(m3, 5, 7)

		for i := 0; i < 3; i++ { // 經過 3 回合
			assert.False(t, character.checkStateIfExpired())

			assert.Equal(t, 1, character.beforeTakeTurn())

			if i == 0 { // 第一回合進行攻擊
				character.attack()
				assert.Equal(t, 0, m.monstersCount()) //怪物全死亡
			}

			character.roundEnd()
		}

		assert.True(t, character.checkStateIfExpired())
		assert.Equal(t, NewTeleport(character).String(), character.state.String())
	})

	t.Run("怪物", func(t *testing.T) {
		m := NewMap(10, NewGame())
		monster := NewMonster(m)
		monster.applyState(NewErupting(monster))
		m.putRoleAt(monster, 2, 3)
		character := NewCharacter(m)
		m.putRoleAt(character, 5, 7) // 不在 monster 原本的攻擊範圍內

		for i := 0; i < 3; i++ { // 經過 3 回合
			assert.False(t, monster.checkStateIfExpired())

			assert.Equal(t, 1, monster.beforeTakeTurn())

			if i == 0 { // 第一回合進行攻擊
				monster.attack()
				assert.Equal(t, character.fullHP-50, character.hp) // character 原本滿血，被攻擊後 HP-50
			}

			monster.roundEnd()
		}

		assert.True(t, monster.checkStateIfExpired())
		assert.Equal(t, NewTeleport(monster).String(), monster.state.String())
	})
}

func TestTeleportStateScenario(t *testing.T) {
	m := NewMap(10, NewGame())
	character := NewCharacter(m)
	character.setDirection(enum.Down)
	character.applyState(NewTeleport(character))
	originalRow := 2
	originalCol := 3
	m.putRoleAt(character, originalRow, originalCol)

	// 經過一回合
	assert.False(t, character.checkStateIfExpired())
	assert.Equal(t, 1, character.beforeTakeTurn())
	character.roundEnd()

	assert.True(t, character.checkStateIfExpired())
	assert.True(t, character.row != originalRow && character.col != originalCol)
	assert.Equal(t, NewNormal(character).String(), character.state.String())
	assert.Equal(t, 1, len(m.roles))
}
