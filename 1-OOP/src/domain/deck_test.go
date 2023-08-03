package domain

import (
	"bigger-or-smaller-game/domain/enum"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDrawCard(t *testing.T) {

	t.Run("when draw 4 cards from deck with 52 cards then deck left only 48 cards", func(t *testing.T) {
		deck := NewDeck()
		expectedLeft := 51

		for i := 0; i < 4; i++ {
			drewCard := deck.DrawCard()
			assert.Equal(t, expectedLeft, len(deck.getCards()))

			findCard := false
			for _, card := range deck.getCards() {
				if card.Rank == drewCard.Rank && card.Suit == drewCard.Suit {
					findCard = true
				}
			}
			assert.False(t, findCard)
			expectedLeft = expectedLeft - 1
		}
	})

	t.Run("when draw 4 cards from deck with 4 cards then deck left no more cards", func(t *testing.T) {
		deck := &Deck{
			cards: []*Card{
				NewCard(enum.Night, enum.Club),
				NewCard(enum.Ten, enum.Club),
				NewCard(enum.J, enum.Club),
				NewCard(enum.Q, enum.Club),
			},
		}
		expectedLeft := 0

		for i := 0; i < 4; i++ {
			deck.DrawCard()
		}

		assert.Equal(t, expectedLeft, len(deck.getCards()))
	})
}
