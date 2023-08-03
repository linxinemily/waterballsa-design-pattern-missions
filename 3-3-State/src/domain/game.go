package domain

import (
	"math/rand"
	"time"
)

type Game struct {
	round int
	m     *Map
}

func NewGame() *Game {
	game := &Game{
		round: 1,
	}
	game.m = NewMap(10, game)
	return game
}

func (game *Game) Start() {

	game.init()

	for game.m.monstersCount() > 0 {
		game.takeRound()
	}

}

func (game *Game) takeRound() {
	// 檢查各個角色狀態是否到期
	for _, role := range game.m.roles {
		role.checkStateIfExpired()
	}

	for _, role := range game.m.roles {
		// 執行不同狀態的特殊行為

		takeTurnTimes := role.beforeTakeTurn()

		for i := 0; i < takeTurnTimes; i++ {
			role.takeTurn()
		}

		role.roundEnd()
	}

	game.round += 1
}

func (game *Game) init() {
	/** 生成寶物 **/
	treasureGenerators := []*TreasureGenerator{
		NewTreasureGenerator(&TreasureType{name: "無敵星星 (Super Star)", rate: "1/10", stateGenerator: StateInvincibleGenerator{}}),
		NewTreasureGenerator(&TreasureType{name: "毒藥 (Poison)", rate: "1/4", stateGenerator: StatePoisonedGenerator{}}),
		NewTreasureGenerator(&TreasureType{name: "加速藥水 (Accelerating Potion)", rate: "1/5", stateGenerator: StateAcceleratedGenerator{}}),
		NewTreasureGenerator(&TreasureType{name: "補血罐 (Healing Potion)", rate: "3/20", stateGenerator: StateHealingGenerator{}}),
		NewTreasureGenerator(&TreasureType{name: "惡魔果實 (Devil Fruit)", rate: "1/10", stateGenerator: StateOrderlessGenerator{}}),
		NewTreasureGenerator(&TreasureType{name: "王者之印 (King's Rock)", rate: "1/10", stateGenerator: StateStockpileGenerator{}}),
		NewTreasureGenerator(&TreasureType{name: "任意門 (Dokodemo Door)", rate: "1/10", stateGenerator: StateTeleportGenerator{}}),
	}

	for i := 0; i < 10; i++ {
		var treasure *Treasure
		for treasure == nil {
			rand.Seed(time.Now().UnixNano())
			treasureGenerator := treasureGenerators[rand.Intn(len(treasureGenerators))]
			treasure = treasureGenerator.generate()
		}
		game.m.putObjectInRandomPosition(treasure)
	}

	/** 生成主角 **/
	character := NewCharacter(game.m)
	game.m.putRoleInRandomPosition(character)
	//game.roles = append(game.roles, character)

	/** 生成怪物 **/
	for i := 0; i < 5; i++ {
		monster := NewMonster(game.m)
		game.m.putRoleInRandomPosition(monster)
		//game.roles = append(game.roles, monster)
	}

	/** 生成障礙物 **/
	for i := 0; i < 3; i++ {
		game.m.putObjectInRandomPosition(NewObstacle())
	}

	character.printState()
}
