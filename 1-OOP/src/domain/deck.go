package domain

import (
	"bigger-or-smaller-game/domain/enum"
	"math/rand"
	"time"
)

type Deck struct {
	cards []*Card
}

func NewDeck() (d *Deck) {
	cards := make([]*Card, 52)
	var count int
	for suit := enum.Suit(0); suit < enum.Spade+1; suit++ {
		for rank := enum.Rank(0); rank < enum.A+1; rank++ {
			cards[count] = NewCard(rank, suit)
			count++
		}
	}

	return &Deck{
		cards: cards,
	}
}

func (d *Deck) getCards() []*Card {
	return d.cards
}

func (d *Deck) DrawCard() *Card {
	// check cards length
	if len(d.cards) <= 0 {
		return nil
	}

	var randNum int

	if len(d.cards) > 1 {
		// randomly give a card from the deck
		rand.Seed(time.Now().UnixNano())
		min := 0
		max := len(d.cards) - 1
		randNum = min + rand.Intn(max-min)
	}

	card := d.cards[randNum]

	d.removeCardFromDeck(randNum)

	return card
}

func (d *Deck) removeCardFromDeck(i int) *Card {
	removed := d.cards[i]
	d.cards[i] = d.cards[len(d.cards)-1]
	d.cards = d.cards[:len(d.cards)-1]
	return removed
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}
