package shopping_recommend_list

import (
	"math/rand"
	algo "mypj/genetic_algo"
)

// Product 產品屬性
type Product struct {
	price    float32
	weight   float32
	category byte
}

var products = []Product{
	{price: 100, weight: 1, category: 'A'},
	{price: 200, weight: 3, category: 'A'},
	{price: 150, weight: 5, category: 'B'},
	{price: 300, weight: 4, category: 'B'},
	{price: 180, weight: 6, category: 'C'},
	{price: 250, weight: 7, category: 'C'},
}

// 客戶喜好度
var preference = map[byte]float32{
	'A': 0.8,
	'B': 0.6,
	'C': 0.2,
}

const (
	BagCapacity = 10
	Budget      = 700
)

type ShoppingRecommendListGeneFactory struct{}

func (factory *ShoppingRecommendListGeneFactory) Create(index int) algo.Genes {
	return algo.Genes(rand.Intn(int(BagCapacity / products[index].weight)))
}

type ShoppingRecommendListIndividualFactory struct{}

func (factory *ShoppingRecommendListIndividualFactory) CalculateFitness(chromosome []algo.Genes) float64 {
	totalPrice := 0.0
	for i, genes := range chromosome {
		totalPrice += float64(genes) * float64(products[i].price)
	}

	totalWeight := 0.0
	for i, genes := range chromosome {
		totalWeight += float64(genes) * float64(products[i].weight)
	}

	if totalPrice > float64(Budget) || totalWeight > float64(BagCapacity) {
		return 0.0
	}

	totalPreference := 0.0
	for i, genes := range chromosome {
		totalPreference += float64(genes) * float64(preference[products[i].category])
	}

	return totalPreference
}

type ShoppingRecommendListPopulationFactory struct{}

func (factory *ShoppingRecommendListPopulationFactory) ChromosomeLen() int {
	return len(products)
}
