package domain

import (
	"big2/domain/enum"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCardPatternFullHouseComparison(t *testing.T) {

	t.Run("第一張開頭為2", func(t *testing.T) {
		cardPattern1 := NewStraightCardPattern([]*Card{
			{Rank: enum.Two, Suit: enum.Heart},
			{Rank: enum.Three, Suit: enum.Heart},
			{Rank: enum.Five, Suit: enum.Heart},
			{Rank: enum.Four, Suit: enum.Heart},
			{Rank: enum.Six, Suit: enum.Heart},
		})

		cardPattern2 := NewStraightCardPattern([]*Card{
			{Rank: enum.Ten, Suit: enum.Heart},
			{Rank: enum.J, Suit: enum.Heart},
			{Rank: enum.Q, Suit: enum.Heart},
			{Rank: enum.K, Suit: enum.Heart},
			{Rank: enum.A, Suit: enum.Heart},
		})
		assert.True(t, cardPattern1.isBiggerThan(cardPattern2))
	})

	t.Run("兩張都是2開頭", func(t *testing.T) {
		cardPattern1 := NewStraightCardPattern([]*Card{
			{Rank: enum.Two, Suit: enum.Heart},
			{Rank: enum.Three, Suit: enum.Heart},
			{Rank: enum.Five, Suit: enum.Heart},
			{Rank: enum.Four, Suit: enum.Heart},
			{Rank: enum.Six, Suit: enum.Heart},
		})

		cardPattern2 := NewStraightCardPattern([]*Card{
			{Rank: enum.Two, Suit: enum.Spade},
			{Rank: enum.Three, Suit: enum.Spade},
			{Rank: enum.Five, Suit: enum.Spade},
			{Rank: enum.Four, Suit: enum.Spade},
			{Rank: enum.Six, Suit: enum.Spade},
		})
		assert.False(t, cardPattern1.isBiggerThan(cardPattern2))
	})

	t.Run("第二張開頭為2", func(t *testing.T) {
		cardPattern1 := NewStraightCardPattern([]*Card{
			{Rank: enum.K, Suit: enum.Heart},
			{Rank: enum.A, Suit: enum.Heart},
			{Rank: enum.Two, Suit: enum.Heart},
			{Rank: enum.Three, Suit: enum.Heart},
			{Rank: enum.Four, Suit: enum.Heart},
		})

		cardPattern2 := NewStraightCardPattern([]*Card{
			{Rank: enum.Two, Suit: enum.Spade},
			{Rank: enum.Three, Suit: enum.Spade},
			{Rank: enum.Five, Suit: enum.Spade},
			{Rank: enum.Four, Suit: enum.Spade},
			{Rank: enum.Six, Suit: enum.Spade},
		})
		assert.False(t, cardPattern1.isBiggerThan(cardPattern2))
	})

	t.Run("兩張開頭都不是2", func(t *testing.T) {
		cardPattern1 := NewStraightCardPattern([]*Card{
			{Rank: enum.Six, Suit: enum.Heart},
			{Rank: enum.Seven, Suit: enum.Heart},
			{Rank: enum.Eight, Suit: enum.Heart},
			{Rank: enum.Night, Suit: enum.Heart},
			{Rank: enum.Ten, Suit: enum.Heart},
		})

		cardPattern2 := NewStraightCardPattern([]*Card{
			{Rank: enum.Five, Suit: enum.Spade},
			{Rank: enum.Six, Suit: enum.Spade},
			{Rank: enum.Seven, Suit: enum.Spade},
			{Rank: enum.Eight, Suit: enum.Spade},
			{Rank: enum.Night, Suit: enum.Spade},
		})
		assert.True(t, cardPattern1.isBiggerThan(cardPattern2))
	})
}
