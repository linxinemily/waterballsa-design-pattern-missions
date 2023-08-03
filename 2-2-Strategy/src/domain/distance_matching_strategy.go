package domain

import (
	"math"
	"sort"
)

type DistanceMatchingStrategy struct{}

func NewDistanceMatchingStrategy() *DistanceMatchingStrategy {
	return &DistanceMatchingStrategy{}
}

func (d *DistanceMatchingStrategy) Match(i Individual, candidates []Individual) Individual {
	return d.getSortedCandidates(i, candidates)[0]
}

func (d *DistanceMatchingStrategy) getSortedCandidates(i Individual, candidates []Individual) []Individual {
	ix, iy := i.GetXY()
	s := ComparWithCandidatesDistance{
		candidates: candidates,
		targetX:    ix,
		targetY:    iy,
	}
	sort.Sort(s)

	return s.candidates
}

type ComparWithCandidatesDistance struct {
	candidates []Individual
	targetX    int
	targetY    int
}

func (w ComparWithCandidatesDistance) Len() int {
	return len(w.candidates)
}

func (w ComparWithCandidatesDistance) Less(i, j int) bool {
	// target
	tx, ty := w.targetX, w.targetY

	ix, iy := w.candidates[i].GetXY()

	//i 和 target 的距離
	distanceWithi := math.Sqrt(math.Pow(float64(tx-ix), 2) + math.Pow(float64(ty-iy), 2))

	jx, jy := w.candidates[j].GetXY()
	//j 和 target 的距離
	distanceWithj := math.Sqrt(math.Pow(float64(tx-jx), 2) + math.Pow(float64(ty-jy), 2))

	return distanceWithi < distanceWithj
}

func (w ComparWithCandidatesDistance) Swap(i, j int) {
	w.candidates[i], w.candidates[j] = w.candidates[j], w.candidates[i]
}
