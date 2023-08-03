package domain

type FullHouseCardPattern struct {
	*AbstractCardPattern
}

func NewFullHouseCardPattern(cards []*Card) *FullHouseCardPattern {
	return &FullHouseCardPattern{
		AbstractCardPattern: NewAbstractCardPattern(cards),
	}
}

func (pattern *FullHouseCardPattern) isBiggerThan(another CardPattern) bool {
	pattern2, ok := another.(*FullHouseCardPattern)
	if !ok {
		return false
	}

	return pattern.Cards()[0].isBiggerThan(pattern2.Cards()[0])
}

func (pattern *FullHouseCardPattern) Name() string {
	return "葫蘆"
}
