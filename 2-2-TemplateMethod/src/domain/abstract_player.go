package domain

import (
	"errors"
)

type Player[C Card[C]] interface {
	NameSelf() (name string)
	Show() (card *C)
	AddCardIntoHand(card C)
	GetName() (name string)
	SetName(string)
	GetHand() []C
	SetHand([]C)
	SetGame(any)
}

type AbstractPlayer[C Card[C], P Player[C]] struct {
	name     string
	hand     []C
	NameSelf func() string
	game     *AbstractGame[C, P]
}

func (p *AbstractPlayer[C, P]) GetName() string {
	return p.name
}

func (p *AbstractPlayer[C, P]) SetName(name string) {
	p.name = name
}

func (p *AbstractPlayer[C, P]) SetGame(game any) {
	other := game.(*AbstractGame[C, P])
	p.game = other
}

func (p *AbstractPlayer[C, P]) AddCardIntoHand(card C) {
	p.hand = append(p.hand, card)
}

func (p *AbstractPlayer[C, P]) SetHand(cards []C) {
	p.hand = cards
}

func (p *AbstractPlayer[C, P]) GetHand() []C {
	return p.hand
}

func (p *AbstractPlayer[C, P]) getCardFromHand(i int) (*C, error) {
	if i < 0 || i > len(p.hand)-1 {
		return nil, errors.New("invalid index")
	}
	return &p.hand[i], nil
}

func (p *AbstractPlayer[C, P]) removeCardFromHand(i int) (*C, error) {
	if i < 0 || i > len(p.hand)-1 {
		return nil, errors.New("invalid index")
	}
	card := p.hand[i]
	p.hand[i] = p.hand[len(p.hand)-1]
	p.hand = p.hand[:len(p.hand)-1]
	return &card, nil
}
