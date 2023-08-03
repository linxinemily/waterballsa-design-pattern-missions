package domain

import (
	"bigger-or-smaller-game/domain/enum"
)

type Card struct {
	Rank enum.Rank
	Suit enum.Suit
}

func NewCard(rank enum.Rank, suit enum.Suit) (c *Card) {
	return &Card{
		Rank: rank,
		Suit: suit,
	}
}

func (c *Card) CompareTo(card *Card) int {

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
