package domain

import uno_enum "bigger-or-smaller-game/domain/enum/uno"

type UnoCard struct {
	Color  uno_enum.Color
	Number uno_enum.Number
}

func NewUnoCard(color uno_enum.Color, number uno_enum.Number) (c *UnoCard) {
	return &UnoCard{
		Color:  color,
		Number: number,
	}
}

func (c UnoCard) CompareTo(card UnoCard) int {
	return 0
}
