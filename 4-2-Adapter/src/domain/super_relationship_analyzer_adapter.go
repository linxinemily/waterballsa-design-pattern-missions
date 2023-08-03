package domain

import (
	"strings"
)

type SuperRelationshipAnalyzerAdapter struct {
	adaptee ISuperRelationshipAnalyzer
	people  map[string]bool
}

func NewSuperRelationshipAnalyzerAdapter(adaptee ISuperRelationshipAnalyzer) *SuperRelationshipAnalyzerAdapter {
	return &SuperRelationshipAnalyzerAdapter{adaptee, make(map[string]bool, 0)}
}

// Parse
// example input:
// A: B C D
// B: A D E
// C: A E G K M
// D: A B K P
// E: B C J K L
// F: Z
func (adapter *SuperRelationshipAnalyzerAdapter) Parse(script string) IRelationshipGraph {
	// turn to another script format
	relationshipMap := make(map[string][]string)
	var newScript strings.Builder
	script = strings.TrimSpace(script)
	lines := strings.Split(script, "\n")

	for i, line := range lines {
		newLine := strings.Split(line, ": ")

		key := strings.TrimSpace(newLine[0])
		adapter.people[key] = true

		if _, exists := relationshipMap[key]; !exists {
			relationshipMap[key] = make([]string, 0)
		}

		vals := strings.Split(newLine[1], " ")

		for j, val := range vals {
			adapter.people[val] = true
			relationshipMap[key] = append(relationshipMap[key], val)
			var newLine strings.Builder
			newLine.WriteString(key)
			newLine.WriteString(" -- ")
			newLine.WriteString(val)
			isLastLine := i == len(line)-1 && j == len(vals)-1
			if !isLastLine {
				newLine.WriteString("\n")
			}
			newScript.WriteString(newLine.String())
		}
	}
	adapter.adaptee.Init(newScript.String())
	return NewBasicGraph(relationshipMap, adapter.people)
}

func (adapter *SuperRelationshipAnalyzerAdapter) GetMutualFriends(name1 string, name2 string) []string {
	mutualFriends := make([]string, 0)
	for person := range adapter.people {
		if adapter.adaptee.IsMutualFriend(person, name1, name2) {
			mutualFriends = append(mutualFriends, person)
		}
	}
	return mutualFriends
}
