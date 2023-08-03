package domain

import (
	"collision-detecting/utils"
	"fmt"
)

type WaterAndWaterCollision struct {
	*AbstractCollision
}

func NewWaterAndWaterCollision(next *ICollision) *WaterAndWaterCollision {
	return &WaterAndWaterCollision{
		AbstractCollision: NewAbstractCollision(next),
	}
}

func (wwc *WaterAndWaterCollision) match(x1 int, x2 int) bool {
	s1 := wwc.world.getSpriteInPosition(x1)
	s2 := wwc.world.getSpriteInPosition(x2)

	return utils.IsSameType(s1, &Water{}) && utils.IsSameType(s2, &Water{})
}

func (wwc *WaterAndWaterCollision) doHandling(x1 int, x2 int) {
	fmt.Printf("water and water collision, cannot move x1 to x2\n")
}
