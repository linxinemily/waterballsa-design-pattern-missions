package domain

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type HumanPlayer struct {
	*AbstractPlayer
}

func NewHumanPlayer(id int, makeCardPatternHandler *IMakeCardPatternHandler) *HumanPlayer {
	return &HumanPlayer{
		AbstractPlayer: NewAbstractPlayer(id, makeCardPatternHandler),
	}
}

func (p *HumanPlayer) getCardsFromUserInput() []*Card {
	var res []*Card
	for {
		if res != nil {
			break
		}

		p.printHand()

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		indexesOfCards := strings.Split(scanner.Text(), " ")

		if indexesOfCards[0] == "-1" {
			fmt.Println("-1 break")
			break
		}

		var cards []*Card
		for i := 0; i < len(indexesOfCards); i++ {
			intVar, err := strconv.Atoi(indexesOfCards[i])
			if err != nil {
				continue
			}
			cards = append(cards, p.hand[intVar])
		}
		res = cards
	}

	return res
}

func (p *HumanPlayer) nameSelf() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter player name:")
	scanner.Scan()
	name := scanner.Text()
	p.name = name
}

func (p *HumanPlayer) printHand() {
	for i := range p.hand {
		fmt.Printf("%d    ", i)
	}
	fmt.Println()
	for _, c := range p.hand {
		fmt.Printf("%s[%s] ", c.Suit, c.Rank)
	}
	fmt.Println()
}
