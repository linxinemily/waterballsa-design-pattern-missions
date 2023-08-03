package domain

import (
	"fmt"
)

type Game[C Card[C], P Player[C]] interface {
	init()
	initPlayers() []P
	playerDrawCards()
	hasNextRound() bool
	takeRound()
	end()
	Start()
}

func NewAbstractGame[C Card[C], P Player[C]]() *AbstractGame[C, P] {
	abstractGame := &AbstractGame[C, P]{}
	// template method default implementation
	abstractGame.shouldBreakDrawCards = shouldBreakDrawCards[C]
	abstractGame.round = 1

	return abstractGame
}

type AbstractGame[C Card[C], P Player[C]] struct {
	round                int
	deck                 *Deck[C]
	players              []P
	initPlayers          func() []P
	shouldBreakDrawCards func(deck Deck[C], count int) bool
	hasNextRound         func() bool
	takeRound            func()
	end                  func()
	stack                []*C
}

func (g *AbstractGame[C, P]) Start() {
	g.init()

	g.playerDrawCards()

	for {
		if !g.hasNextRound() {
			break
		}
		g.takeRound()
	}

	g.end()
}

func (g *AbstractGame[C, P]) init() {

	players := g.initPlayers()

	for _, p := range players {
		name := p.NameSelf()
		fmt.Printf("Player %s is added. \n", name)
		g.players = append(g.players, p)

		p.SetGame(g)
	}

	g.deck.Shuffle()
}

func (g *AbstractGame[C, P]) playerDrawCards() {
	count := 0
	for {
		if g.shouldBreakDrawCards(*g.deck, count) {
			break
		}
		card := g.deck.DrawCard()
		p := g.players[count%4]
		p.AddCardIntoHand(*card)
		count += 1
	}
}

func shouldBreakDrawCards[C Card[C]](deck Deck[C], count int) bool {
	return len(deck.getCards()) == 0
}

func (g *AbstractGame[C, P]) GetTopCardFromStack() *C {
	if len(g.stack) == 0 {
		return nil
	}
	return g.stack[len(g.stack)-1]
}
