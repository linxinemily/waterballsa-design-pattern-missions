package domain

import (
	"fmt"
	"math/rand"
	"time"
)

type AIPlayer struct {
	AbstractPlayer
}

func NewAIPlayer() (p *AIPlayer) {
	return &AIPlayer{
		AbstractPlayer: NewAbstractPlayer(),
	}
}

func (p *AIPlayer) NameSelf() string {
	rand.Seed(time.Now().UnixNano())
	p.name = fmt.Sprintf("AI Player %d", rand.Intn(999999))
	return p.name
}

func (p *AIPlayer) ToUseExchangeChance() bool {
	rand.Seed(time.Now().UnixNano())
	res := rand.Intn(9)

	if res < 2 {
		return true
	} else {
		return false
	}
}

func (p *AIPlayer) ChoosePlayerForExchange() (player Player) {
	rand.Seed(time.Now().UnixNano())
	randIndex := rand.Intn(2)

	var playersExceptSelf []Player

	for _, otherPlayer := range p.game.players {
		if otherPlayer != p {
			playersExceptSelf = append(playersExceptSelf, otherPlayer)
		}
	}

	return playersExceptSelf[randIndex]
}

func (p *AIPlayer) Show() *Card {
	rand.Seed(time.Now().UnixNano())

	var i int
	if len(p.hand) > 1 {
		i = rand.Intn(len(p.hand) - 1)
	}

	removed, _ := p.removeCardFromHand(i)

	return removed
}
