package domain

import (
	"C3M3H1/domain/enum"
	"fmt"
	"math/rand"
	"time"
)

type Monster struct {
	*AbstractRole
}

func NewMonster(m *Map) *Monster {
	monster := &Monster{
		NewAbstractRole(m, "🦖", 1),
	}
	monster.applyState(NewNormal(monster))
	return monster
}

func (monster *Monster) takeTurn() {
	if monster.hp <= 0 { //如果這輪已經遭受攻擊死亡，就不再繼續
		return
	}
	//如果主角沒有位於怪物的攻擊範圍之內的話，怪物將會自主決定要往哪一個方向移動一格
	if monster.getValidEnemies() == nil {
		rand.Seed(time.Now().UnixNano())
		stateGenerator, err := monster.moveTo(monster, enum.AllDirections[rand.Intn(len(enum.AllDirections))])
		if err == nil && stateGenerator != nil {
			fmt.Println("怪物得到禮物", stateGenerator.generate(monster).String())
			state := stateGenerator.generate(monster)
			fmt.Printf("怪物套用 %s 狀態", state)
			monster.applyState(state)
		}
	} else { //否則怪物會站在原地攻擊主角
		fmt.Println("怪物攻擊主角")
		monster.attack()
	}
	fmt.Printf("怪物 HP: %d, State: %s\n", monster.hp, monster.state)
}

func (monster *Monster) getValidEnemies() []Role {
	// 取得攻擊範圍內的主角
	for _, direction := range enum.AllDirections {
		row, col := enum.GetMovementByDirection(monster.row, monster.col, direction)
		if char := monster.getCharacterAt(row, col); char != nil { // top
			return []Role{char}
		}
	}
	return nil
}

func (monster *Monster) getCharacterAt(row, col int) *Character {
	if obj, err := monster.getMap().getObjectAt(row, col); err == nil && obj != nil {
		character, isCharacter := obj.(*Character)
		if isCharacter {
			return character
		}
	}
	return nil
}

func (monster *Monster) afterAttacked() {
	// 死亡，直接從地圖移除
	monster.hp = 0
	monster.m.removeObject(monster)
}

func (monster *Monster) getAllEnemies() []Role {
	var characters []Role
	for _, role := range monster.m.roles {
		if character, isCharacter := role.(*Character); isCharacter {
			characters = append(characters, character)
		}
	}
	return characters
}
