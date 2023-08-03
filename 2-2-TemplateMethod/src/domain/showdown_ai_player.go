package domain

import (
	"fmt"
	"math/rand"
	"time"
)

type ShowdownAIPlayer struct {
	AbstractShowdownPlayer
}

func NewShowdownAIPlayer() (p *ShowdownAIPlayer) {
	return &ShowdownAIPlayer{
		AbstractShowdownPlayer: NewAbstractShowdownPlayer(),
	}
}

func (p *ShowdownAIPlayer) NameSelf() string {
	rand.Seed(time.Now().UnixNano())
	p.name = fmt.Sprintf("AI Player %d", rand.Intn(999999))
	return p.name
}

func (p *ShowdownAIPlayer) ToUseExchangeChance() bool {
	rand.Seed(time.Now().UnixNano())
	res := rand.Intn(9)

	if res < 2 {
		return true
	} else {
		return false
	}
}

func (p *ShowdownAIPlayer) ChoosePlayerForExchange() (showdownPlayer ShowdownPlayer) {
	rand.Seed(time.Now().UnixNano())
	randIndex := rand.Intn(2)

	var playersExceptSelf []ShowdownPlayer
	for _, otherPlayer := range p.game.players {
		if otherPlayer.GetName() != p.name {
			playersExceptSelf = append(playersExceptSelf, otherPlayer)
		}
	}

	return playersExceptSelf[randIndex]
}

func (p *ShowdownAIPlayer) Show() *ShowdownCard {
	rand.Seed(time.Now().UnixNano())

	var i int
	if len(p.hand) > 1 {
		i = rand.Intn(len(p.hand) - 1)
	}

	removed, _ := p.removeCardFromHand(i)

	return removed
}
