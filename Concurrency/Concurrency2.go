package Concurrency

import (
	"fmt"
	"sync"
)

func generateData() []int {
	data := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		data[i] = i
	}
	return data
}

func worker(data []int, results chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for _, num := range data {
		sum += num
	}

	results <- sum
}

func ExerciseTwo() {
	wg := sync.WaitGroup{}
	numWorkers := 5
	data := generateData()

	results := make(chan int, numWorkers)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(data, results, &wg)
	}

	wg.Wait()
	close(results)

	sum := 0
	for res := range results {
		sum += res
	}

	fmt.Println("Sum:", sum)
}
