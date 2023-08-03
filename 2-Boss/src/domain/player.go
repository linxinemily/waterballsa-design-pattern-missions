package domain

import (
	"fmt"
	"sort"
)

type IPlayer struct {
	Player
}

func (p *IPlayer) takeTurn(turn int) CardPattern {
	fmt.Printf("輪到%s了\n", p.Name())
	canPass := true
	if turn == 0 { // 新回合首輪玩家不能喊 pass
		canPass = false
	}

	for {
		cards := p.getCardsFromUserInput()
		if len(cards) == 0 { // pass
			if canPass {
				return nil
			} else {
				fmt.Println("cannot pass")
			}
		} else {
			cardPattern := p.play(cards)
			if cardPattern != nil {
				return cardPattern
			}
		}
	}
}

type Player interface {
	play(cards []*Card) CardPattern
	nameSelf()
	addCardIntoHand(card *Card)
	Hand() []*Card
	Id() int
	Name() string
	getCardsFromUserInput() []*Card
	removeCardFromHand(card *Card) *Card
}

type AbstractPlayer struct {
	CardPatternHandler *IMakeCardPatternHandler
	name               string
	id                 int
	hand               []*Card
}

func NewAbstractPlayer(id int, makeCardPatternHandler *IMakeCardPatternHandler) *AbstractPlayer {
	return &AbstractPlayer{
		id:                 id,
		CardPatternHandler: makeCardPatternHandler,
	}
}

func (p *AbstractPlayer) play(cards []*Card) CardPattern {
	return p.CardPatternHandler.handle(cards)
}

func (p *AbstractPlayer) addCardIntoHand(card *Card) {
	p.hand = append(p.hand, card)
	sort.Sort(CardSlice(p.hand))
}

func (p *AbstractPlayer) Hand() []*Card {
	return p.hand
}

func (p *AbstractPlayer) Id() int {
	return p.id
}

func (p *AbstractPlayer) Name() string {
	return p.name
}

func (p *AbstractPlayer) removeCardFromHandByIdx(i int) (*Card, error) {
	card := p.hand[i]
	p.hand[i] = p.hand[len(p.hand)-1]
	p.hand = p.hand[:len(p.hand)-1]
	sort.Sort(CardSlice(p.hand))
	return card, nil
}

func (p *AbstractPlayer) removeCardFromHand(card *Card) *Card {
	for i, handCard := range p.hand {
		if handCard.isEuqalTo(card) {
			r, _ := p.removeCardFromHandByIdx(i)
			return r
		}
	}
	return nil
}
