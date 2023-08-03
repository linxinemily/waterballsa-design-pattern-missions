package domain

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type ShowdownHumanPlayer struct {
	AbstractShowdownPlayer
}

func NewShowdownHumanPlayer() (p *ShowdownHumanPlayer) {
	return &ShowdownHumanPlayer{
		AbstractShowdownPlayer: NewAbstractShowdownPlayer(),
	}
}

func (p *ShowdownHumanPlayer) NameSelf() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter player name:")
	scanner.Scan()
	name := scanner.Text()
	p.name = name
	return name
}

func (p *ShowdownHumanPlayer) ToUseExchangeChance() bool {

	var res bool

	fmt.Println("Do you want to exchange hand this round?(y/N)")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	answer := scanner.Text()
	if answer == "y" {
		res = true
	}

	return res
}

func (p *ShowdownHumanPlayer) ChoosePlayerForExchange() (showdownPlayer ShowdownPlayer) {

	var playersExceptSelf []ShowdownPlayer

	for _, otherPlayer := range p.game.players {
		if otherPlayer.GetName() != p.name {
			playersExceptSelf = append(playersExceptSelf, otherPlayer)
		}
	}

	first := true
	var err error
	var intVar int

	for {
		if err == nil && !first {
			break
		}

		fmt.Println("Choose a player which you want to exchange hand (enter index of player)")

		for i, p := range playersExceptSelf {
			fmt.Printf("[%d] %s \n", i, p.GetName())
		}
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		intVar, err = strconv.Atoi(scanner.Text())

		if intVar < 0 || intVar > len(playersExceptSelf)-1 {
			err = errors.New("index out of range")
		}

		first = false
	}

	return playersExceptSelf[intVar]
}

func (p *ShowdownHumanPlayer) Show() (card *ShowdownCard) {

	first := true
	var err error
	var intVar int
	var removed *ShowdownCard

	for {
		if err == nil && !first {
			break
		}

		fmt.Println("Your hand:")
		for i, c := range p.hand {
			fmt.Printf("[%d] rank %s, suit %s \n", i, c.Rank, c.Suit)
		}
		fmt.Println("Choose a card to show, enter index of the card:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		intVar, err = strconv.Atoi(scanner.Text())
		removed, err = p.removeCardFromHand(intVar)

		first = false
	}

	fmt.Printf("show card: rank %s, suit %s \n", removed.Rank, removed.Suit)

	return removed
}
