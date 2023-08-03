package main

import (
	"fmt"
	"matching-system/domain"
	"matching-system/domain/enum"
)

func main() {

	i1 := domain.NewIndividual(1, enum.Male, 20, "hi", "唱歌, 聽音樂, 跳舞, 看電影", "(1,2)")
	i2 := domain.NewIndividual(2, enum.Female, 22, "hi", "睡覺, 打電動", "(4,8)")
	i3 := domain.NewIndividual(3, enum.Male, 23, "hi", "購物, 唱歌, 聽音樂, 看電影", "(8,16)")
	i4 := domain.NewIndividual(4, enum.Female, 21, "hi", "睡覺, 美食, 旅遊, 聽音樂", "(16,32)")
	i5 := domain.NewIndividual(5, enum.Male, 20, "hi", "旅遊, 美食, 爬山, 看電影", "(32,64)")

	// st := domain.NewDistanceMatchingStrategy()
	// ms := domain.NewMatchingSystem([]domain.Individual{*i1, *i2, *i3, *i4, *i5}, st)
	// res := ms.Match(*i1)
	// fmt.Println(res)

	// st := domain.NewDistanceMatchingStrategy()
	// ms := domain.NewMatchingSystem([]domain.Individual{*i1, *i2, *i3, *i4, *i5}, domain.NewReverseMatchingStrategy(st))
	// res := ms.Match(*i1)
	// fmt.Println(res)

	// hst := domain.NewHabitMatchingStrategy()
	// ms := domain.NewMatchingSystem([]domain.Individual{*i1, *i2, *i3, *i4, *i5}, hst)
	// res := ms.Match(*i1)
	// fmt.Println(res)

	hst := domain.NewHabitMatchingStrategy()
	ms := domain.NewMatchingSystem([]domain.Individual{*i1, *i2, *i3, *i4, *i5}, domain.NewReverseMatchingStrategy(hst))
	res := ms.Match(*i1)
	fmt.Println(res)

	// type Person struct {
	// 	Name string
	// }

	// slice1 := []Person{{Name: "A"}, {Name: "B"}, {Name: "C"}}
	// slice2 := slice1

	// // fmt.Printf("Addr of first element: %p\n", &slice1[0])
	// // fmt.Printf("Addr of first element: %p\n", &slice2[0])

	// slice2[0] = slice2[2]

	// fmt.Printf("Addr of first element: %p\n", &slice1[0])
	// fmt.Println(slice1[0])
	// fmt.Printf("Addr of first element: %p\n", &slice2[0])
	// fmt.Println(slice2[0])

	// slice2 = slice2[:2]

	// // fmt.Printf("Addr of first element: %p\n", &slice1[0])
	// // fmt.Printf("Addr of first element: %p\n", &slice2[0])
	// fmt.Println(slice1)
	// fmt.Println(slice2)
}
