package genetic_algo

import (
	"math/rand"
	"time"
)

type GeneticAlgo struct {
	maxIteration      int
	individualFactory *IndividualFactoryWrapper
	genesFactory      GenesFactory
}

func NewGeneticAlgo(maxIteration int, individualFactory *IndividualFactoryWrapper, genesFactory GenesFactory) *GeneticAlgo {
	return &GeneticAlgo{maxIteration, individualFactory, genesFactory}
}

func (algo *GeneticAlgo) Evolve(p Population) *Individual {
	rand.Seed(time.Now().UnixNano())

	currentPopulation := p

	for i := 0; i < algo.maxIteration; i++ {
		newPopulation := make(Population, 0)
		for j := 0; j < len(currentPopulation)/2; j++ {
			parent1, parent2 := algo.selection(currentPopulation)
			child1, child2 := algo.crossover(parent1, parent2)

			child1 = algo.mutation(child1)
			child2 = algo.mutation(child2)

			newPopulation = append(newPopulation, child1, child2)
		}

		currentPopulation = newPopulation

		if terminationCondition() {
			break
		}
	}

	return findBestIndividual(currentPopulation)
}

func findBestIndividual(population Population) *Individual {
	bestIndividual := population[0]
	for _, individual := range population {
		if individual.GetFitness() > bestIndividual.GetFitness() {
			bestIndividual = individual
		}
	}
	return bestIndividual
}

func terminationCondition() bool {
	return false
}

func (algo *GeneticAlgo) mutation(i *Individual) *Individual {
	chromosome := i.GetChromosome()
	for j := 0; j < len(chromosome); j++ {
		if rand.Float64() < 0.1 {
			chromosome[j] = algo.genesFactory.Create(j)
		}
	}
	return algo.individualFactory.Create(chromosome)
}

func (algo *GeneticAlgo) crossover(p1 *Individual, p2 *Individual) (c1 *Individual, c2 *Individual) {
	crossoverPoint := rand.Intn(len(p1.GetChromosome()))
	c1 = algo.individualFactory.Create(append(p1.GetChromosome()[:crossoverPoint], p2.GetChromosome()[crossoverPoint:]...))
	c2 = algo.individualFactory.Create(append(p2.GetChromosome()[:crossoverPoint], p1.GetChromosome()[crossoverPoint:]...))
	return c1, c2
}
func (algo *GeneticAlgo) selection(population Population) (i1 *Individual, i2 *Individual) {
	currentPopulation := population
	selectedNum := 2
	selected := make(Population, selectedNum)
	tournamentSize := 2

	for i := 0; i < selectedNum; i++ {
		tournament := make(Population, tournamentSize)

		// 隨機選取參與錦標賽的個體
		for j := 0; j < tournamentSize; j++ {
			randomIndex := rand.Intn(len(currentPopulation))
			tournament[j] = currentPopulation[randomIndex]
		}

		// 找出錦標賽中適應度最高的個體
		winner := findWinner(tournament)

		selected[i] = winner
	}

	return selected[0], selected[1]
}

func findWinner(tournament Population) *Individual {
	best := tournament[0]
	for i := 1; i < len(tournament); i++ {
		if tournament[i].GetFitness() > best.GetFitness() {
			best = tournament[i]
		}
	}
	return best
}
