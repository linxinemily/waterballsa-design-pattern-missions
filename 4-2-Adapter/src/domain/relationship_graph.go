package domain

type IRelationshipGraph interface {
	HasConnection(name1 string, name2 string) bool
}
