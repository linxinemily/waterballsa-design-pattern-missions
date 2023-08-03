package domain

import (
	"bigger-or-smaller-game/domain/enum/showdown"
)

func NewShowdownDeck() (d *Deck[ShowdownCard]) {
	cards := make([]*ShowdownCard, 52)
	var count int
	for suit := showdown_enum.Suit(0); suit < showdown_enum.Spade+1; suit++ {
		for rank := showdown_enum.Rank(0); rank < showdown_enum.A+1; rank++ {
			cards[count] = NewShowdownCard(rank, suit)
			count++
		}
	}

	return &Deck[ShowdownCard]{
		cards: cards,
	}
}
