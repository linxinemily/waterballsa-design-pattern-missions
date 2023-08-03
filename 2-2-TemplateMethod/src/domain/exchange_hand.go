package domain

import "fmt"

type ExchangeHand struct {
	shouldRollbackRound int
	p1                  ShowdownPlayer
	p2                  ShowdownPlayer
}

func NewExchangeHand(shouldRollbackRound int, p1 ShowdownPlayer, p2 ShowdownPlayer) (e *ExchangeHand) {
	return &ExchangeHand{
		shouldRollbackRound,
		p1,
		p2,
	}
}

func (e *ExchangeHand) Rollback() {
	e.Exchange()
	fmt.Printf("Roll back hand from %s to %s \n", e.p2.GetName(), e.p1.GetName())
}

func (e *ExchangeHand) Exchange() {
	temp := e.p1.GetHand()
	e.p1.SetHand(e.p2.GetHand())
	e.p2.SetHand(temp)
}

func (e *ExchangeHand) haveToRollback(currentRound int) bool {
	return currentRound == e.shouldRollbackRound
}
