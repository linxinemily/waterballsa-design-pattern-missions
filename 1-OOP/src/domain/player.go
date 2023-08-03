package domain

import (
	"errors"
)

type Player interface {
	NameSelf() (name string)
	ToUseExchangeChance() bool
	ChoosePlayerForExchange() (player Player)
	Show() (card *Card)
	AddCardIntoHand(card *Card)
	GetName() (name string)
	SetExchangeHand(*ExchangeHand)
	CanUseExchangeHand() bool
	GetPoints() int
	SetPoints(int)
	SetGame(*Game)
	GetHand() []*Card
	SetHand([]*Card)
	GetExchangeHand() *ExchangeHand
}

type AbstractPlayer struct {
	exchangeChance int
	points         int
	name           string
	game           *Game
	hand           []*Card
	exchangeHand   *ExchangeHand
}

func NewAbstractPlayer() (pa AbstractPlayer) {
	return AbstractPlayer{
		exchangeChance: 1,
		points:         0,
	}
}

func (p *AbstractPlayer) AddCardIntoHand(card *Card) {
	p.hand = append(p.hand, card)
}

func (p *AbstractPlayer) GetName() string {
	return p.name
}

func (p *AbstractPlayer) SetExchangeHand(exchangeHand *ExchangeHand) {
	p.exchangeHand = exchangeHand
	p.exchangeChance = p.exchangeChance - 1
	p.exchangeHand.Exchange()
}

func (p *AbstractPlayer) GetExchangeHand() *ExchangeHand {
	return p.exchangeHand
}

func (p *AbstractPlayer) CanUseExchangeHand() bool {
	return p.exchangeChance > 0
}

func (p *AbstractPlayer) removeCardFromHand(i int) (*Card, error) {
	if i < 0 || i > len(p.hand)-1 {
		return nil, errors.New("invalid index")
	}
	card := p.hand[i]
	p.hand[i] = p.hand[len(p.hand)-1]
	p.hand = p.hand[:len(p.hand)-1]
	return card, nil
}

func (p *AbstractPlayer) GetPoints() int {
	return p.points
}

func (p *AbstractPlayer) SetPoints(points int) {
	p.points = points
}

func (p *AbstractPlayer) SetGame(game *Game) {
	p.game = game
}

func (p *AbstractPlayer) SetHand(cards []*Card) {
	p.hand = cards
}

func (p *AbstractPlayer) GetHand() []*Card {
	return p.hand
}
