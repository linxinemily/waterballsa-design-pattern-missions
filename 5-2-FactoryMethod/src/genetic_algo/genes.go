package genetic_algo

type Genes rune

type GenesFactory interface {
	Create(index int) Genes
}
