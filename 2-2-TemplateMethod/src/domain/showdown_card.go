package domain

import (
	"bigger-or-smaller-game/domain/enum/showdown"
)

type ShowdownCard struct {
	Rank showdown_enum.Rank
	Suit showdown_enum.Suit
}

func NewShowdownCard(rank showdown_enum.Rank, suit showdown_enum.Suit) (c *ShowdownCard) {
	return &ShowdownCard{
		Rank: rank,
		Suit: suit,
	}
}

func (c ShowdownCard) CompareTo(card ShowdownCard) int {

	if c.Rank > card.Rank {
		return 1
	}

	if c.Rank < card.Rank {
		return -1
	}

	// rank is equal, compare suit
	if c.Suit > card.Suit {
		return 1
	}

	if c.Suit < card.Suit {
		return -1
	}

	return 0
}
