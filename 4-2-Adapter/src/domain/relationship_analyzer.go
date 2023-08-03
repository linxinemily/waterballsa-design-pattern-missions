package domain

type RelationshipAnalyzer interface {
	Parse(script string)
	GetMutualFriends(name1 string, name2 string) []string
}
