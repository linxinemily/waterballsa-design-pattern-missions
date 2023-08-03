package domain

import (
	"fmt"
)

type ShowdownGame struct {
	AbstractGame[ShowdownCard, ShowdownPlayer]
	totalRound int
}

func NewShowdownGame() *ShowdownGame {

	showdownGame := &ShowdownGame{
		totalRound: 13,
	}

	showdownGame.AbstractGame = *NewAbstractGame[ShowdownCard, ShowdownPlayer]()

	showdownGame.AbstractGame.initPlayers = showdownGame.initPlayers
	showdownGame.AbstractGame.hasNextRound = showdownGame.hasNextRound
	showdownGame.AbstractGame.takeRound = showdownGame.takeRound
	showdownGame.AbstractGame.end = showdownGame.end
	showdownGame.deck = NewShowdownDeck()

	return showdownGame
}

func (g *ShowdownGame) initPlayers() []ShowdownPlayer {
	p1 := NewShowdownAIPlayer()
	p2 := NewShowdownAIPlayer()
	p3 := NewShowdownAIPlayer()
	p4 := NewShowdownHumanPlayer()

	players := []ShowdownPlayer{p1, p2, p3, p4}

	return players
}

func (g *ShowdownGame) takeRound() {
	res := make(map[string]*ShowdownCard)
	var winner ShowdownPlayer
	var maxCard *ShowdownCard

	fmt.Println("--------------------")
	fmt.Printf("Round %d \n", g.round)
	fmt.Println("--------------------")

	for _, p := range g.players {
		fmt.Printf("It's %s's turn \n", p.GetName())

		// check if there are exchange hands should be roll back
		exchangeHand := p.GetExchangeHand()
		if exchangeHand != nil && exchangeHand.haveToRollback(g.round) {
			exchangeHand.Rollback()
		}

		if p.CanUseExchangeHand() && p.ToUseExchangeChance() {
			exchangePlayer := p.ChoosePlayerForExchange()
			fmt.Printf("Player %s exchange hand with %s \n", p.GetName(), exchangePlayer.GetName())

			exchangeHand := NewExchangeHand(g.round+3, p, exchangePlayer)
			p.SetExchangeHand(exchangeHand)
		}

		card := p.Show()
		res[p.GetName()] = card
		if card == nil {
			fmt.Printf("Player %s has no more hand \n", p.GetName())
			continue
		}

		if maxCard == nil {
			maxCard = &ShowdownCard{} // dummy
		}
		if card.CompareTo(*maxCard) > 0 {
			maxCard = card
			winner = p
		}
	}

	// print result
	for name, card := range res {
		if card != nil {
			fmt.Printf("Player: %s shows card, rank: %s, suit: %s \n", name, card.Rank, card.Suit)
		}
	}
	fmt.Printf("The winner of round %d is %s \n", g.round, winner.GetName())
	winner.SetPoints(winner.GetPoints() + 1)
	g.round += 1
}

func (g *ShowdownGame) end() {
	// calculate result and declare the winner
	var winner ShowdownPlayer
	var highestPoints int
	for _, p := range g.players {
		if p.GetPoints() > highestPoints {
			highestPoints = p.GetPoints()
			winner = p
		}
		fmt.Printf("Player: %s's points: %d\n", p.GetName(), p.GetPoints())
	}
	fmt.Printf("The winner is: %s \n", winner.GetName())

}

func (g *ShowdownGame) hasNextRound() bool {
	return g.round <= g.totalRound
}
