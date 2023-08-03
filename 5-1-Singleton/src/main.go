package main

import (
	"5-1/domain"
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	totalTasks := 100
	wg.Add(totalTasks)

	for i := 0; i < totalTasks; i++ {
		go func() {
			models := domain.GetModels()
			model, err := models.CreateModel("Reflection")
			if err != nil {
				panic(err)
			}
			var arr [1000]float32
			for i := 0; i < 1000; i++ {
				arr[i] = 1.0
			}
			r := model.Count(arr)

			fmt.Println(r)
			wg.Done()
		}()
	}

	wg.Wait()
}
