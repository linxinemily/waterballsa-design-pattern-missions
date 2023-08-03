package domain

import (
	"collision-detecting/utils"
	"fmt"
)

type HeroAndHeroCollision struct {
	*AbstractCollision
}

func NewHeroAndHeroCollision(next *ICollision) *HeroAndHeroCollision {
	return &HeroAndHeroCollision{
		AbstractCollision: NewAbstractCollision(next),
	}
}

func (wwc *HeroAndHeroCollision) match(x1 int, x2 int) bool {
	s1 := wwc.world.getSpriteInPosition(x1)
	s2 := wwc.world.getSpriteInPosition(x2)

	return utils.IsSameType(s1, &Hero{}) && utils.IsSameType(s2, &Hero{})
}

func (wwc *HeroAndHeroCollision) doHandling(x1 int, x2 int) {
	fmt.Println("hero and hero collision, can not move x1 to x2")
}
