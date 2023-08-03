package domain

import (
	"big2/domain/enum"
	"fmt"
	"strings"
)

type CardPattern interface {
	isBiggerThan(another CardPattern) bool
	printCards() string
	containsClub3() bool
	Cards() []*Card
	Name() string
}

type AbstractCardPattern struct {
	cards []*Card
}

func NewAbstractCardPattern(cards []*Card) *AbstractCardPattern {
	return &AbstractCardPattern{
		cards: cards,
	}
}

func (cardPattern *AbstractCardPattern) printCards() string {
	var sb strings.Builder
	for _, card := range cardPattern.cards {
		sb.WriteString(fmt.Sprintf("%s[%s] ", card.Suit, card.Rank))
	}
	return fmt.Sprintf(sb.String())
}

func (cardPattern *AbstractCardPattern) containsClub3() bool {
	for _, card := range cardPattern.cards {
		if card.Rank == enum.Three && card.Suit == enum.Club {
			return true
		}
	}

	return false
}

func (cardPattern *AbstractCardPattern) Cards() []*Card {
	return cardPattern.cards
}
