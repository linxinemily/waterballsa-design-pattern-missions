package factory_schedule

import (
	"math/rand"
	algo "mypj/genetic_algo"
)

type Product struct {
	name   byte
	time   int
	demand int
}

var products = []Product{
	{name: 'A', time: 2, demand: 100},
	{name: 'B', time: 3, demand: 200},
	{name: 'C', time: 4, demand: 300},
}

type ScheduleGeneFactory struct{}

func (factory *ScheduleGeneFactory) Create(index int) algo.Genes {
	return algo.Genes(rand.Intn(products[index].demand))
}

type ScheduleIndividualFactory struct{}

func (factory *ScheduleIndividualFactory) CalculateFitness(chromosome []algo.Genes) float64 {
	fitness := 0.0
	for i, p := range products {
		fitness += float64(chromosome[i]) * float64(p.time)
	}
	return -fitness
}

type SchedulePopulationFactory struct{}

func (factory *SchedulePopulationFactory) ChromosomeLen() int {
	return len(products)
}
