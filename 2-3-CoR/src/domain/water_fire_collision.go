package domain

import (
	"collision-detecting/utils"
	"fmt"
)

type WaterAndFireCollision struct {
	*AbstractCollision
}

func NewWaterAndFireCollision(next *ICollision) *WaterAndFireCollision {
	return &WaterAndFireCollision{
		AbstractCollision: NewAbstractCollision(next),
	}
}

func (wwc *WaterAndFireCollision) match(x1 int, x2 int) bool {
	s1 := wwc.world.getSpriteInPosition(x1)
	s2 := wwc.world.getSpriteInPosition(x2)
	return utils.IsSameType(s1, &Water{}) && utils.IsSameType(s2, &Fire{}) || utils.IsSameType(s1, &Fire{}) && utils.IsSameType(s2, &Water{})
}

func (wwc *WaterAndFireCollision) doHandling(x1 int, x2 int) {
	fmt.Println("water and fire collision")
	wwc.world.removeSpriteInPosition(x1)
	wwc.world.removeSpriteInPosition(x2)
}
