package domain

type PairCardPattern struct {
	*AbstractCardPattern
}

func NewPairCardPattern(cards []*Card) *PairCardPattern {
	return &PairCardPattern{
		AbstractCardPattern: NewAbstractCardPattern(cards),
	}
}

func (pattern *PairCardPattern) isBiggerThan(another CardPattern) bool {
	pattern2, ok := another.(*PairCardPattern)
	if !ok {
		return false
	}

	return pattern.Cards()[0].isBiggerThan(pattern2.Cards()[0])
}

func (pattern *PairCardPattern) Name() string {
	return "對子"
}
