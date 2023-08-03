package domain

type ReverseMatchingStrategy struct {
	ms MatchingStrategy
}

func NewReverseMatchingStrategy(ms MatchingStrategy) MatchingStrategy {
	return &ReverseMatchingStrategy{
		ms,
	}
}

func (d *ReverseMatchingStrategy) Match(i Individual, candidates []Individual) Individual {
	return d.getSortedCandidates(i, candidates)[len(candidates)-1]
}

func (d *ReverseMatchingStrategy) getSortedCandidates(i Individual, candidates []Individual) []Individual {
	return d.ms.getSortedCandidates(i, candidates)
}
