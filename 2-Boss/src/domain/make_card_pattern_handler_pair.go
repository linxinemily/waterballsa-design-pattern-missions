package domain

import "sort"

type MakePairCardPatternHandler struct {
	*AbstractMakeCardPatternHandler
}

func NewMakePairCardPatternHandler(next *IMakeCardPatternHandler) *MakePairCardPatternHandler {
	return &MakePairCardPatternHandler{
		AbstractMakeCardPatternHandler: NewAbstractHandler(next),
	}
}

func (handler *MakePairCardPatternHandler) match(cards []*Card) (cardPattern CardPattern, ok bool) {
	ok = len(cards) == 2 && cards[0].Rank == cards[1].Rank
	sort.Sort(sort.Reverse(CardSlice(cards))) //降冪排序
	return NewPairCardPattern(cards), ok
}
