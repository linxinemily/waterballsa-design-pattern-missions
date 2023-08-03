package domain

import (
	"collision-detecting/utils"
	"fmt"
)

type HeroAndWaterCollision struct {
	*AbstractCollision
}

func NewHeroAndWaterCollision(next *ICollision) *HeroAndWaterCollision {
	return &HeroAndWaterCollision{
		AbstractCollision: NewAbstractCollision(next),
	}
}

func (wwc *HeroAndWaterCollision) match(x1 int, x2 int) bool {
	s1 := wwc.world.getSpriteInPosition(x1)
	s2 := wwc.world.getSpriteInPosition(x2)

	return utils.IsSameType(s1, &Hero{}) && utils.IsSameType(s2, &Water{}) || utils.IsSameType(s1, &Water{}) && utils.IsSameType(s2, &Hero{})
}

func (wwc *HeroAndWaterCollision) doHandling(x1 int, x2 int) {
	fmt.Println("hero and water collision")
	s1 := wwc.world.getSpriteInPosition(x1)
	s2 := wwc.world.getSpriteInPosition(x2)

	var hero *Hero
	if (utils.IsSameType(s1, &Hero{}) && utils.IsSameType(s2, &Water{})) {
		hero, _ = s1.(*Hero)
		hero.addHP(10)
		// move hero to x2
		wwc.world.setSpriteInPosition(x2, hero)
		wwc.world.removeSpriteInPosition(x1)
	} else if (utils.IsSameType(s1, &Water{}) && utils.IsSameType(s2, &Hero{})) {
		hero, _ = s2.(*Hero)
		hero.addHP(10)
		wwc.world.removeSpriteInPosition(x1) // just remove water
	}
}
