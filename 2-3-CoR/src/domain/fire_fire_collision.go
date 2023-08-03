package domain

import (
	"collision-detecting/utils"
	"fmt"
)

type FireAndFireCollision struct {
	*AbstractCollision
}

func NewFireAndFireCollision(next *ICollision) *FireAndFireCollision {
	return &FireAndFireCollision{
		AbstractCollision: NewAbstractCollision(next),
	}
}

func (wwc *FireAndFireCollision) match(x1 int, x2 int) bool {
	s1 := wwc.world.getSpriteInPosition(x1)
	s2 := wwc.world.getSpriteInPosition(x2)

	return utils.IsSameType(s1, &Fire{}) && utils.IsSameType(s2, &Fire{})
}

func (wwc *FireAndFireCollision) doHandling(x1 int, x2 int) {
	fmt.Printf("fire and fire collision, cannot move x1 to x2\n")
}
