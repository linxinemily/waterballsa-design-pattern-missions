package domain

import (
	"big2/domain/enum"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatchFullHouseWithValidData(t *testing.T) {
	type MatchFullHouseData struct {
		input  []*Card
		expect []*Card
	}
	validInputs := []MatchFullHouseData{
		{
			input: []*Card{
				{Rank: enum.Night, Suit: enum.Heart},
				{Rank: enum.A, Suit: enum.Diamond},
				{Rank: enum.Night, Suit: enum.Diamond},
				{Rank: enum.Night, Suit: enum.Spade},
				{Rank: enum.A, Suit: enum.Heart},
			},
			expect: []*Card{
				{Rank: enum.Night, Suit: enum.Heart},
				{Rank: enum.Night, Suit: enum.Diamond},
				{Rank: enum.Night, Suit: enum.Spade},
				{Rank: enum.A, Suit: enum.Diamond},
				{Rank: enum.A, Suit: enum.Heart},
			},
		},
		{
			input: []*Card{
				{Rank: enum.Two, Suit: enum.Spade},
				{Rank: enum.Two, Suit: enum.Heart},
				{Rank: enum.Five, Suit: enum.Heart},
				{Rank: enum.Five, Suit: enum.Diamond},
				{Rank: enum.Five, Suit: enum.Spade},
			},
			expect: []*Card{
				{Rank: enum.Five, Suit: enum.Heart},
				{Rank: enum.Five, Suit: enum.Diamond},
				{Rank: enum.Five, Suit: enum.Spade},
				{Rank: enum.Two, Suit: enum.Spade},
				{Rank: enum.Two, Suit: enum.Heart},
			},
		},
		{
			input: []*Card{
				{Rank: enum.Three, Suit: enum.Heart},
				{Rank: enum.Three, Suit: enum.Diamond},
				{Rank: enum.Three, Suit: enum.Spade},
				{Rank: enum.Four, Suit: enum.Spade},
				{Rank: enum.Four, Suit: enum.Heart},
			},
			expect: []*Card{
				{Rank: enum.Three, Suit: enum.Heart},
				{Rank: enum.Three, Suit: enum.Diamond},
				{Rank: enum.Three, Suit: enum.Spade},
				{Rank: enum.Four, Suit: enum.Spade},
				{Rank: enum.Four, Suit: enum.Heart},
			},
		},
	}

	for i := 0; i < len(validInputs); i++ {
		t.Run(fmt.Sprintf("valid inputs %v", i), func(t *testing.T) {
			output, ok := NewMakeFullHouseCardPatternHandler(nil).match(validInputs[i].input)
			assert.True(t, ok, output)
			for j := 0; j < 5; j++ {
				assert.True(t, output.Cards()[j].isEuqalTo(validInputs[i].expect[j]))
			}
		})
	}
}
func TestMatchFullHouseWithInValidData(t *testing.T) {
	t.Run("invalid input", func(t *testing.T) {
		cards := []*Card{
			{
				Rank: enum.Three,
				Suit: enum.Heart,
			},
			{
				Rank: enum.A,
				Suit: enum.Diamond,
			},
			{
				Rank: enum.Night,
				Suit: enum.Diamond,
			},
			{
				Rank: enum.Night,
				Suit: enum.Spade,
			},
			{
				Rank: enum.A,
				Suit: enum.Heart,
			},
		}

		_, ok := NewMakeFullHouseCardPatternHandler(nil).match(cards)
		assert.False(t, ok)
	})
}
