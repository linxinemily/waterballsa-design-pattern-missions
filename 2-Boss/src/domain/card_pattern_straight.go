package domain

import "big2/domain/enum"

type StraightCardPattern struct {
	*AbstractCardPattern
}

func NewStraightCardPattern(cards []*Card) *StraightCardPattern {
	return &StraightCardPattern{
		AbstractCardPattern: NewAbstractCardPattern(cards),
	}
}

func (pattern *StraightCardPattern) isBiggerThan(another CardPattern) bool {
	pattern2, ok := another.(*StraightCardPattern)
	if !ok {
		return false
	}

	// 比較兩副牌中最大的牌(已排好的最後一張牌)
	// 例外：23456 為最大
	if pattern.Cards()[0].Rank == enum.Two && pattern2.Cards()[0].Rank == enum.Two { // 兩張都是2開頭
		return pattern.Cards()[0].isBiggerThan(pattern2.Cards()[0]) // 比較第一張牌誰大
	} else if pattern.Cards()[0].Rank == enum.Two { // pattern 第一張是2開頭
		return true // pattern 一定比 pattern2 大
	} else if pattern2.Cards()[0].Rank == enum.Two { // pattern2 第一張是2開頭
		return false // pattern2 一定比 pattern 大
	}

	// 兩張都不是2開頭，直接比最後一張牌
	return pattern.Cards()[4].isBiggerThan(pattern2.Cards()[4])
}

func (pattern *StraightCardPattern) Name() string {
	return "順子"
}
