package domain

import (
	"sort"
)

type MakeFullHouseCardPatternHandler struct {
	*AbstractMakeCardPatternHandler
}

func NewMakeFullHouseCardPatternHandler(next *IMakeCardPatternHandler) *MakeFullHouseCardPatternHandler {
	return &MakeFullHouseCardPatternHandler{
		AbstractMakeCardPatternHandler: NewAbstractHandler(next),
	}
}

func (handler *MakeFullHouseCardPatternHandler) match(cards []*Card) (cardPattern CardPattern, ok bool) {
	if len(cards) != 5 {
		return nil, false
	}

	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Rank > cards[j].Rank
	})

	// 2 + 3 組合 => 需對調（3張要在前面）
	if cards[0].Rank == cards[1].Rank && cards[2].Rank == cards[3].Rank && cards[3].Rank == cards[4].Rank {
		mergeCards := make([]*Card, 0)
		mergeCards = append(mergeCards, cards[2:]...)
		mergeCards = append(mergeCards, cards[:2]...)
		return NewFullHouseCardPattern(mergeCards), true
	}

	// 3 + 2 組合 => 直接回傳
	if cards[0].Rank == cards[1].Rank && cards[1].Rank == cards[2].Rank && cards[3].Rank == cards[4].Rank {
		return NewFullHouseCardPattern(cards), true
	}

	return nil, false
}
