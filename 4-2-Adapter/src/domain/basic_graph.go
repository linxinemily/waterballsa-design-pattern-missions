package domain

import (
	"github.com/yourbasic/graph"
	"golang.org/x/exp/slices"
)

type BasicGraph struct {
	g           *graph.Mutable
	personIdMap map[string]int
	components  [][]int
}

func NewBasicGraph(relationshipMap map[string][]string, people map[string]bool) *BasicGraph {

	g := graph.New(len(people))

	personIdMap := make(map[string]int)

	counter := 0
	for person := range people {
		personIdMap[person] = counter
		counter += 1
	}

	for person, friends := range relationshipMap {
		for _, friend := range friends {
			g.Add(personIdMap[person], personIdMap[friend])
		}
	}

	components := graph.Components(g)
	return &BasicGraph{g, personIdMap, components}
}

func (bg BasicGraph) HasConnection(name1 string, name2 string) bool {
	for _, component := range bg.components {
		if slices.Contains(component, bg.personIdMap[name1]) && slices.Contains(component, bg.personIdMap[name2]) {
			return true
		}
	}
	return false
}
