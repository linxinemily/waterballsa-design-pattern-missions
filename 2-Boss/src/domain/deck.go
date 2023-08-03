package domain

import (
	"big2/domain/enum"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

type Deck struct {
	cards []*Card
}

func NewDeck() (d *Deck) {
	cards := make([]*Card, 52)
	var count int
	for suit := enum.Suit(0); suit < enum.Suit(3)+1; suit++ {
		for rank := enum.Rank(0); rank < enum.Rank(12)+1; rank++ {
			cards[count] = NewCard(rank, suit)
			count++
		}
	}
	return &Deck{
		cards: cards,
	}
}

// NewDeckFromCardsInput for testing
func NewDeckFromCardsInput(inputString string) (d *Deck) {
	res := strings.Split(inputString, " ")

	cards := make([]*Card, 52)

	for i, str := range res {
		regex := *regexp.MustCompile("([A-z])\\[(.*)\\]")
		res := regex.FindAllStringSubmatch(str, -1)
		suitStr := res[0][1]
		rankStr := res[0][2]

		cards[i] = NewCard(enum.RankStringToVal(rankStr), enum.SuitStringToVal(suitStr))
	}

	return &Deck{
		cards: cards,
	}
}

func (d *Deck) deal() *Card {
	// check cards length
	cardsLen := len(d.cards)

	if cardsLen <= 0 {
		return nil
	}
	card := d.cards[cardsLen-1]
	d.cards = d.cards[:cardsLen-1]

	return card
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

func (d *Deck) Cards() []*Card {
	return d.cards
}
