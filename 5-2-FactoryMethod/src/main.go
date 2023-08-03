package main

import (
	"fmt"
	"mypj/factory_schedule"
	"mypj/genetic_algo"
)

func main() {

	// 工廠排程問題
	genesFactory := &factory_schedule.ScheduleGeneFactory{}
	individualFactory := &genetic_algo.IndividualFactoryWrapper{
		IndividualFactory: &factory_schedule.ScheduleIndividualFactory{},
	}
	algo := genetic_algo.NewGeneticAlgo(20, individualFactory, genesFactory)
	populationFactory := &genetic_algo.PopulationFactoryWrapper{
		PopulationFactory: &factory_schedule.SchedulePopulationFactory{},
	}
	population := populationFactory.InitRandomly(20, individualFactory, genesFactory)

	// 購物推薦清單問題
	//genesFactory := &shopping_recommend_list.ShoppingRecommendListGeneFactory{}
	//individualFactory := &genetic_algo.IndividualFactoryWrapper{
	//	IndividualFactory: &shopping_recommend_list.ShoppingRecommendListIndividualFactory{},
	//}
	//algo := genetic_algo.NewGeneticAlgo(20, individualFactory, genesFactory)
	//
	//populationFactory := &genetic_algo.PopulationFactoryWrapper{
	//	PopulationFactory: &shopping_recommend_list.ShoppingRecommendListPopulationFactory{},
	//}
	//population := populationFactory.InitRandomly(20, individualFactory, genesFactory)

	i := algo.Evolve(population)

	fmt.Println(i.GetChromosome(), i.GetFitness())
}
