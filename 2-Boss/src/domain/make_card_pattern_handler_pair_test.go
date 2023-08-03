package domain

import (
	"big2/domain/enum"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatchPairWithValidInput(t *testing.T) {
	type MatchPairData struct {
		input  []*Card
		expect []*Card
	}
	validInputs := []MatchPairData{
		{
			input: []*Card{
				{Rank: enum.Four, Suit: enum.Diamond},
				{Rank: enum.Four, Suit: enum.Heart},
			},
			expect: []*Card{
				{Rank: enum.Four, Suit: enum.Heart},
				{Rank: enum.Four, Suit: enum.Diamond},
			},
		},
		{
			input: []*Card{
				{Rank: enum.Ten, Suit: enum.Spade},
				{Rank: enum.Ten, Suit: enum.Club},
			},
			expect: []*Card{
				{Rank: enum.Ten, Suit: enum.Spade},
				{Rank: enum.Ten, Suit: enum.Club},
			},
		},
	}

	for i := 0; i < len(validInputs); i++ {
		t.Run(fmt.Sprintf("valid inputs %v", i), func(t *testing.T) {
			output, ok := NewMakePairCardPatternHandler(nil).match(validInputs[i].input)
			assert.True(t, ok, output)
			for j := 0; j < len(validInputs[i].input); j++ {
				assert.True(t, output.Cards()[j].isEuqalTo(validInputs[i].expect[j]))
			}
		})
	}
}

func TestMatchPairWithInvalidInput(t *testing.T) {
	invalidInputs := []*Card{
		{Rank: enum.Three, Suit: enum.Heart},
		{Rank: enum.A, Suit: enum.Diamond},
	}

	t.Run("invalid input", func(t *testing.T) {
		_, ok := NewMakePairCardPatternHandler(nil).match(invalidInputs)
		assert.False(t, ok)
	})
}
