package main

import (
	"c3-boss/domain"
)

func main() {
	rpg := domain.NewRPG()

	hero := rpg.CreateHero("英雄", 500, 500, 100)
	hero.SetSkills([]*domain.SkillImpl{{domain.NewBasicSkill(hero)}})

	waterTA := rpg.CreateAI("WaterTA", 200, 200, 70)
	waterTA.SetSkills([]*domain.SkillImpl{{domain.NewBasicSkill(hero)}})

	troop1 := rpg.CreateTroop(hero, waterTA)

	slime1 := rpg.CreateAI("Slime1", 200, 90, 50)
	slime1.SetSkills([]*domain.SkillImpl{{domain.NewBasicSkill(slime1)}})

	slime2 := rpg.CreateAI("Slime2", 200, 90, 50)
	slime2.SetSkills([]*domain.SkillImpl{{domain.NewBasicSkill(slime2)}})

	slime3 := rpg.CreateAI("Slime3", 200, 9000, 50)
	slime3.SetSkills([]*domain.SkillImpl{{domain.NewBasicSkill(slime3)}})

	troop2 := rpg.CreateTroop(slime1, slime2, slime3)

	rpg.StartBattle(troop1, troop2)
}
