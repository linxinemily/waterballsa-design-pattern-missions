package domain

import (
	"C3M3H1/domain/enum"
)

type IRole struct {
	Role
}

type Role interface {
	MapObject
	attack()
	takeTurn()
	applyState(state *IState)
	checkStateIfExpired() bool
	beforeTakeTurn() int
	afterAttacked()
	getMap() *Map
	getValidEnemies() []Role
	getAllEnemies() []Role
	attacked()
	subtractHP(hp int)
	addHP(hp int)
	isFullHP() bool
	roundEnd()
	getHP() int
	getFullHP() int
}

type AbstractRole struct {
	m *Map
	*AbstractMapObject
	state  *IState
	hp     int
	fullHP int
}

func NewAbstractRole(m *Map, symbol string, hp int) *AbstractRole {
	return &AbstractRole{
		AbstractMapObject: NewAbstractMapObject(symbol),
		m:                 m,
		hp:                hp,
		fullHP:            hp, //滿血
	}
}

func (role *AbstractRole) checkStateIfExpired() bool {
	return role.state.checkStateIfExpired()
}

func (role *AbstractRole) subtractHP(hp int) {
	role.hp -= hp
	if role.hp <= 0 {
		role.m.removeObject(role)
	}
}

func (role *AbstractRole) addHP(hp int) {
	role.hp += hp
	if role.hp >= role.fullHP {
		role.hp = role.fullHP
	}
}

func (role *AbstractRole) getHP() int {
	return role.hp
}

func (role *AbstractRole) getFullHP() int {
	return role.fullHP
}

func (role *AbstractRole) isFullHP() bool {
	return role.hp >= role.fullHP
}

func (role *AbstractRole) getMap() *Map {
	return role.m
}

func (role *AbstractRole) attack() {
	role.state.attack()
}

func (role *AbstractRole) attacked() {
	role.state.attacked()
}

func (role *AbstractRole) beforeTakeTurn() int {
	return role.state.beforeTakeTurn()
}

func (role *AbstractRole) applyState(state *IState) {
	role.state = state
}

func (role *AbstractRole) moveTo(self Role, direction enum.RoleDirection) (StateGenerator, error) {

	updatedRow := role.Row()
	updatedCol := role.Col()

	switch direction {
	case enum.Top:
		updatedRow -= 1
	case enum.Down:
		updatedRow += 1
	case enum.Left:
		updatedCol -= 1
	case enum.Right:
		updatedCol += 1
	}

	// 移動目的位置上有東西
	objectAtTarget, err := role.getMap().getObjectAt(updatedRow, updatedCol)
	if err != nil {
		return nil, err
	}
	var stateGenerator StateGenerator
	if objectAtTarget != nil {
		if treasure, isTreasure := objectAtTarget.(*Treasure); isTreasure {
			stateGenerator = treasure.stateGenerator
			err := role.getMap().moveRoleTo(self, updatedRow, updatedCol)
			if err != nil {
				return nil, err
			}
		}
	} else {
		err := role.getMap().moveRoleTo(self, updatedRow, updatedCol)
		if err != nil {
			return nil, err
		}
	}

	return stateGenerator, nil
}

func (role *AbstractRole) roundEnd() {
	role.state.increaseDuringRound()
}
