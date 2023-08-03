package domain

import (
	"fmt"
	"math/rand"
	"time"
)

type UnoAIPlayer struct {
	AbstractUnoPlayer
}

func NewUnoAIPlayer() (p *UnoAIPlayer) {
	UnoAIPlayer := &UnoAIPlayer{
		AbstractUnoPlayer: NewAbstractUnoPlayer(),
	}
	return UnoAIPlayer
}

func (p *UnoAIPlayer) NameSelf() string {
	rand.Seed(time.Now().UnixNano())
	p.name = fmt.Sprintf("AI Player %d", rand.Intn(999999))
	return p.name
}

func (p *UnoAIPlayer) Show() *UnoCard {
	rand.Seed(time.Now().UnixNano())

	var validateCardsIdx []int
	CardOfStackTop := p.game.GetTopCardFromStack()

	for index, card := range p.hand {
		if p.isValidateCard(&card, CardOfStackTop) {
			validateCardsIdx = append(validateCardsIdx, index)
		}
	}

	var i int
	if len(validateCardsIdx) > 1 {
		i = rand.Intn(len(validateCardsIdx) - 1)
	}

	removed, _ := p.removeCardFromHand(validateCardsIdx[i])

	return removed
}
