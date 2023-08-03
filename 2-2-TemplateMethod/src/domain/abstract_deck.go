package domain

import (
	"math/rand"
	"time"
)

type Deck[T any] struct {
	cards []*T
}

func (d *Deck[T]) DrawCard() *T {
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

func (d *Deck[T]) removeCardFromDeck(i int) *T {
	removed := d.cards[i]
	d.cards[i] = d.cards[len(d.cards)-1]
	d.cards = d.cards[:len(d.cards)-1]
	return removed
}

func (d *Deck[T]) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

func (d *Deck[T]) getCards() []*T {
	return d.cards
}

func (d *Deck[T]) setCards(cards []*T) {
	d.cards = cards
} 