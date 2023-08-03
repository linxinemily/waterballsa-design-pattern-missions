package domain

import (
	"big2/domain/enum"
	"sort"
)

type MakeStraightCardPatternHandler struct {
	*AbstractMakeCardPatternHandler
}

func NewMakeStraightCardPatternHandler(next *IMakeCardPatternHandler) *MakeStraightCardPatternHandler {
	return &MakeStraightCardPatternHandler{
		AbstractMakeCardPatternHandler: NewAbstractHandler(next),
	}
}

func (handler *MakeStraightCardPatternHandler) match(cards []*Card) (cardPattern CardPattern, ok bool) {
	if len(cards) != 5 {
		return nil, false
	}

	// 先依照大小（升冪）排序完，再分組
	// 分組規則：如果中間有斷層(相減大於一)就分成另外一組
	// 最後結果如果第一組的第一個是 Three(0)，且第二組的最後一個是 Two(12)，表示可以接起來 => 為合法牌型
	sort.SliceStable(cards, func(i, j int) bool {
		return cards[i].Rank < cards[j].Rank
	})

	group := []map[int]*Card{
		{},
	}

	for i := 0; i < len(cards)-1; i++ {

		group[len(group)-1][int(cards[i].Rank)] = cards[i]

		if int(cards[i+1].Rank)-int(cards[i].Rank) == 1 {
			group[len(group)-1][int(cards[i+1].Rank)] = cards[i+1]
		} else {
			group = append(group, map[int]*Card{
				int(cards[i+1].Rank): cards[i+1],
			})
		}
	}

	if len(group[0]) == 5 { // 沒碰到尾接頭
		return NewStraightCardPattern(cards), true
	}

	if len(group[1]) > 0 { // 有碰到尾接頭,變成兩組
		var firstGroupKeys []int
		for k := range group[0] {
			firstGroupKeys = append(firstGroupKeys, k)
		}
		sort.Ints(firstGroupKeys)

		var SecondGroupKeys []int
		for k := range group[1] {
			SecondGroupKeys = append(SecondGroupKeys, k)
		}
		sort.Ints(SecondGroupKeys)
		if firstGroupKeys[0] == int(enum.Three) && SecondGroupKeys[len(SecondGroupKeys)-1] == int(enum.Two) {
			var sortedCards []*Card
			// 先接尾巴那組
			// ex:[[3,4][K,A,2]] => 先接 [K,A,2]
			for _, key := range SecondGroupKeys {
				sortedCards = append(sortedCards, group[1][key])
			}
			// 再接頭
			for _, key := range firstGroupKeys {
				sortedCards = append(sortedCards, group[0][key])
			}

			return NewStraightCardPattern(sortedCards), true
		}
	}

	return nil, false
}
