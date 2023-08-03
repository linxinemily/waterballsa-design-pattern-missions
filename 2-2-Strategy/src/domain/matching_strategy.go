package domain

type MatchingStrategy interface {
	Match(i Individual, candidates []Individual) Individual
	getSortedCandidates(i Individual, candidates []Individual) []Individual
}
