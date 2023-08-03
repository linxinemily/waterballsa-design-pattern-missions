package domain

import (
	"sort"
)

type HabitMatchingStrategy struct{}

func NewHabitMatchingStrategy() *HabitMatchingStrategy {
	return &HabitMatchingStrategy{}
}

func (d *HabitMatchingStrategy) Match(i Individual, candidates []Individual) Individual {
	return d.getSortedCandidates(i, candidates)[0]
}

func (d *HabitMatchingStrategy) getSortedCandidates(i Individual, candidates []Individual) []Individual {
	s := ComparWithCandidatesHabitsCount{
		candidates:   candidates,
		targetHabits: i.GetHabits(),
	}

	sort.Sort(s)
	return s.candidates
}

type ComparWithCandidatesHabitsCount struct {
	candidates   []Individual
	targetHabits []string
}

func (c ComparWithCandidatesHabitsCount) Len() int {
	return len(c.candidates)
}

func (c ComparWithCandidatesHabitsCount) Less(i, j int) bool {
	// target
	iHabits := c.candidates[i].GetHabits()
	jHabits := c.candidates[j].GetHabits()

	//i 和 target 興趣重疊的數量
	intersectionWithi := intersection(c.targetHabits, iHabits)

	//j 和 target 興趣重疊的數量
	intersectionWithj := intersection(c.targetHabits, jHabits)

	return len(intersectionWithj) < len(intersectionWithi)
}

func (c ComparWithCandidatesHabitsCount) Swap(i, j int) {
	c.candidates[i], c.candidates[j] = c.candidates[j], c.candidates[i]
}

func intersection(s1, s2 []string) (inter []string) {
	hash := make(map[string]bool)
	for _, e := range s1 {
		hash[e] = true
	}
	for _, e := range s2 {
		if hash[e] {
			inter = append(inter, e)
		}
	}
	return
}
