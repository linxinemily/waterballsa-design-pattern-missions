package domain

import (
	"golang.org/x/exp/slices"
	"strings"
)

type ISuperRelationshipAnalyzer interface {
	Init(script string)
	IsMutualFriend(targetName string, name2 string, name3 string) bool
}

type SuperRelationshipAnalyzer struct {
	relationshipMap map[string][]string
}

func NewSuperRelationshipAnalyzer() *SuperRelationshipAnalyzer {
	return &SuperRelationshipAnalyzer{
		relationshipMap: make(map[string][]string),
	}
}

// Init
//
//	example of script:
//	A -- B
//	A -- C
//	A -- D
//	B -- D
//	B -- E
//	C -- E
//	C -- G
//	C -- K
//	C -- M
//	D -- K
//	D -- P
//	E -- J
//	E -- K
//	E -- L
//	F -- Z
func (analyzer *SuperRelationshipAnalyzer) Init(script string) {
	script = strings.TrimSpace(script)
	lines := strings.Split(script, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		keyVal := strings.Split(line, " -- ")

		if _, exists := analyzer.relationshipMap[keyVal[0]]; !exists {
			analyzer.relationshipMap[keyVal[0]] = make([]string, 0)
		}
		analyzer.relationshipMap[keyVal[0]] = append(analyzer.relationshipMap[keyVal[0]], keyVal[1])
	}
}

func (analyzer *SuperRelationshipAnalyzer) IsMutualFriend(targetName string, name2 string, name3 string) bool {
	// 如果 name2 和 targetName 是好友，name3 和 targetName 是好友，targetName 為 name2 和 name3 的共同好友

	return (slices.Contains(analyzer.relationshipMap[name2], targetName) || slices.Contains(analyzer.relationshipMap[targetName], name2)) &&
		(slices.Contains(analyzer.relationshipMap[name3], targetName) || slices.Contains(analyzer.relationshipMap[targetName], name3))
}
