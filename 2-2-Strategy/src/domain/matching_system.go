package domain

import (
	"golang.org/x/exp/slices"
)

type MatchingSystem struct {
	members []Individual
	ms      MatchingStrategy
}

func NewMatchingSystem(members []Individual, ms MatchingStrategy) *MatchingSystem {
	return &MatchingSystem{
		members,
		ms,
	}
}

func (m *MatchingSystem) Match(i Individual) Individual {
	return m.ms.Match(i, getMembersExceptSelf(i, m.members))
}

func getMembersExceptSelf(i Individual, members []Individual) []Individual {
	exceptIndex := slices.IndexFunc(members, func(m Individual) bool { return m.GetId() == i.GetId() })

	membersExcpetSelf := append([]Individual(nil), members...)
	membersExcpetSelf[exceptIndex] = membersExcpetSelf[len(membersExcpetSelf)-1]
	membersExcpetSelf = membersExcpetSelf[:len(membersExcpetSelf)-1]

	return membersExcpetSelf
}
