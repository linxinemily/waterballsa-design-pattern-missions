package genetic_algo

type Individual struct {
	chromosome []Genes
	fitness    float64
}

func NewIndividual(chromosome []Genes, fitness float64) *Individual {
	return &Individual{chromosome, fitness}
}

func (i *Individual) GetFitness() float64 {
	return i.fitness
}

func (i *Individual) GetChromosome() []Genes {
	return i.chromosome
}

type IndividualFactoryWrapper struct {
	IndividualFactory
}

func (factory *IndividualFactoryWrapper) Create(chromosome []Genes) *Individual {
	return NewIndividual(chromosome, factory.CalculateFitness(chromosome))
}

type IndividualFactory interface {
	CalculateFitness(chromosome []Genes) float64
}
