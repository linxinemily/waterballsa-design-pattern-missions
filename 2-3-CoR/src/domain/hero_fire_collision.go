package domain

import (
	"collision-detecting/utils"
	"fmt"
)

type HeroAndFireCollision struct {
	*AbstractCollision
}

func NewHeroAndFireCollision(next *ICollision) *HeroAndFireCollision {
	return &HeroAndFireCollision{
		AbstractCollision: NewAbstractCollision(next),
	}
}

func (wwc *HeroAndFireCollision) match(x1 int, x2 int) bool {
	s1 := wwc.world.getSpriteInPosition(x1)
	s2 := wwc.world.getSpriteInPosition(x2)

	return utils.IsSameType(s1, &Hero{}) && utils.IsSameType(s2, &Fire{}) || utils.IsSameType(s1, &Fire{}) && utils.IsSameType(s2, &Hero{})
}

func (wwc *HeroAndFireCollision) doHandling(x1 int, x2 int) {
	fmt.Println("hero and fire collision")
	s1 := wwc.world.getSpriteInPosition(x1)
	s2 := wwc.world.getSpriteInPosition(x2)

	var hero *Hero
	if (utils.IsSameType(s1, &Hero{}) && utils.IsSameType(s2, &Fire{})) {
		hero, _ = s1.(*Hero)
		hero.substractHP(10)
		wwc.world.removeSpriteInPosition(x2)
		wwc.world.removeSpriteInPosition(x1)
		if hero.HP <= 0 {
			fmt.Println("hero died")
		} else {
			wwc.world.setSpriteInPosition(x2, hero)
		}
	} else if (utils.IsSameType(s1, &Fire{}) && utils.IsSameType(s2, &Hero{})) {
		hero, _ = s2.(*Hero)
		hero.substractHP(10)
		wwc.world.removeSpriteInPosition(x1)
		if hero.HP <= 0 {
			wwc.world.removeSpriteInPosition(x2)
			fmt.Println("hero died")
		}
	}
}
