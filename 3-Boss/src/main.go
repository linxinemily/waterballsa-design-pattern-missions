package main

import "c3-boss/domain"

func main() {
	rpg := domain.NewRPG()

	hero := rpg.CreateHero("勇者", 100, 100, 10)
	ai1 := rpg.CreateAI("スライム", 50, 0, 5)
	ai2 := rpg.CreateAI("スライム", 50, 0, 5)
	ai3 := rpg.CreateAI("スライム", 50, 0, 5)
	ai4 := rpg.CreateAI("スライム", 50, 0, 5)
	ai5 := rpg.CreateAI("スライム", 50, 0, 5)

	troop1 := rpg.CreateTroop(hero, ai1, ai2)
	troop2 := rpg.CreateTroop(ai3, ai4, ai5)

	rpg.StartBattle(troop1, troop2)
}
