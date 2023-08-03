package domain

type UnoPlayer interface {
	Player[UnoCard]
	hasNoCardCanShow(UnoCard) bool
}

type AbstractUnoPlayer struct {
	AbstractPlayer[UnoCard, UnoPlayer]
	NameSelf func() (name string)
	Show     func() (card *UnoCard)
}

func NewAbstractUnoPlayer() (pa AbstractUnoPlayer) {
	unoPlayer := AbstractUnoPlayer{
		AbstractPlayer: AbstractPlayer[UnoCard, UnoPlayer]{},
	}

	return unoPlayer
}

func (p *AbstractUnoPlayer) hasNoCardCanShow(topCard UnoCard) bool {
	var hasNoCardCanShow = true
	for _, card := range p.GetHand() {
		if p.isValidateCard(&card, &topCard) {
			return false
		}
	}
	return hasNoCardCanShow
}

func (p *AbstractUnoPlayer) isValidateCard(choseCard *UnoCard, topCard *UnoCard) bool {
	return choseCard.Color == topCard.Color || choseCard.Number == topCard.Number
}
