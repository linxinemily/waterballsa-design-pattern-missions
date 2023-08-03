package domain

import "big2/domain/enum"

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

func (card *Card) isEuqalTo(card2 *Card) bool {
	return card.compare(card2) == 0
}

func (card *Card) isBiggerThan(card2 *Card) bool {
	return card.compare(card2) == 1
}

func (card *Card) compare(card2 *Card) int {
	if card.Rank > card2.Rank {
		return 1
	}

	if card.Rank < card2.Rank {
		return -1
	}

	// rank is equal, compare suit
	if card.Suit > card2.Suit {
		return 1
	}

	if card.Suit < card2.Suit {
		return -1
	}

	return 0
}

type CardSlice []*Card

func (cards CardSlice) Len() int           { return len(cards) }
func (cards CardSlice) Less(i, j int) bool { return cards[i].compare(cards[j]) == -1 }
func (cards CardSlice) Swap(i, j int)      { cards[i], cards[j] = cards[j], cards[i] }
