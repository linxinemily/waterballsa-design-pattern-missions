package genetic_algo

type Population []*Individual

type PopulationFactory interface {
	ChromosomeLen() int
}

type PopulationFactoryWrapper struct {
	PopulationFactory
}

func (pf *PopulationFactoryWrapper) InitRandomly(
	nums int,
	individualFactory *IndividualFactoryWrapper,
	geneFactory GenesFactory,
) Population {
	s := make([]*Individual, nums)
	for i := 0; i < nums; i++ {
		chromosome := make([]Genes, pf.ChromosomeLen())
		for j := 0; j < pf.ChromosomeLen(); j++ {
			chromosome[j] = geneFactory.Create(j)
		}
		s[i] = individualFactory.Create(chromosome)
	}
	return s
}
