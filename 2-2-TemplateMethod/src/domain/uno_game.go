package domain

import "fmt"

type UnoGame struct {
	AbstractGame[UnoCard, UnoPlayer]
	winner *UnoPlayer
}

func NewUnoGame() *UnoGame {
	unoGame := &UnoGame{
		AbstractGame: *NewAbstractGame[UnoCard, UnoPlayer](),
	}

	unoGame.AbstractGame.initPlayers = unoGame.initPlayers
	unoGame.AbstractGame.hasNextRound = unoGame.hasNextRound
	unoGame.AbstractGame.takeRound = unoGame.takeRound
	unoGame.AbstractGame.end = unoGame.end

	unoGame.AbstractGame.shouldBreakDrawCards = unoGame.shouldBreakDrawCards
	unoGame.deck = NewUnoDeck()

	return unoGame
}

func (g *UnoGame) initPlayers() []UnoPlayer {
	p1 := NewUnoAIPlayer()
	p2 := NewUnoAIPlayer()
	p3 := NewUnoAIPlayer()
	p4 := NewUnoHumanPlayer()

	players := []UnoPlayer{p1, p2, p3, p4}

	return players
}

func (g *UnoGame) shouldBreakDrawCards(deck Deck[UnoCard], count int) bool {
	return len(deck.getCards())-4*5 <= 0
}

func (g *UnoGame) takeRound() {

	fmt.Println("--------------------")
	fmt.Printf("Round %d \n", g.round)
	fmt.Println("--------------------")

	topCard := g.GetTopCardFromStack()
	if topCard == nil {
		fmt.Println("there are no card in stack, draw a card from deck")
		topCard = g.safeDrawCard()
		g.stack = append(g.stack, topCard)
	}
	fmt.Printf("Top card color: %s, number: %s \n", topCard.Color, topCard.Number)

	for _, p := range g.players {
		fmt.Printf("It's %s's turn \n", p.GetName())

		//玩家沒有可出的牌
		if p.hasNoCardCanShow(*g.GetTopCardFromStack()) {
			// 玩家就必須從牌堆中抽一張牌，如果此時牌堆空了，則會先把檯面上除了最新的牌以外的牌放回牌堆中進行洗牌
			newCard := g.safeDrawCard()
			p.SetHand(append(p.GetHand(), *newCard))

		} else {
			card := p.Show()
			if len(p.GetHand()) == 0 {
				fmt.Printf("Player %s has no more hand.\n", p.GetName())
				g.winner = &p
				return
			} else {
				fmt.Printf("Player: %s shows card, color: %s, number: %s \n", p.GetName(), card.Color, card.Number)
				g.stack = append(g.stack, card)
			}

		}
	}
	g.round += 1

}

func (g *UnoGame) end() {
	w := *g.winner
	fmt.Printf("The winner is %s \n", w.GetName())
}

func (g *UnoGame) hasNextRound() bool {
	return g.winner == nil
}

func (g *UnoGame) safeDrawCard() *UnoCard {
	newCard := g.deck.DrawCard()

	if newCard != nil {
		return newCard
	}
	g.deck.setCards(g.stack[:len(g.stack)-1])
	g.deck.Shuffle()

	return g.deck.DrawCard()
}

