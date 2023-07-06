package ConcurrencyIdiomatic

import (
	"fmt"
	"sync"
	"time"
)

func ConcurrencyIdiomatic3() {
	var intStream []int
	channel1 := make(chan int)
	channel2 := make(chan int)
	channel3 := make(chan int)

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		intStream = append(intStream, i)
	}

	initial := func(channel1 chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("initial")
		for i := range intStream {
			fmt.Println("Initial sending:", i)
			channel1 <- i
			time.Sleep(1)
		}
		close(channel1)
	}

	stageOne := func(channel1 chan int, channel2 chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("Stage One")
		for elem := range channel1 {
			fmt.Println("Stage 1 received", elem)
			channel2 <- elem*2 + 1
		}
		close(channel2)
	}

	stageTwo := func(channel2 chan int, channel3 chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("Stage Two")
		for elem := range channel2 {
			fmt.Println("Stage 2 received", elem)
			if elem%2 != 0 {
				channel3 <- elem
			}
		}
		close(channel3)
	}

	stageThree := func(channel3 chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("Stage Three")
		sum := 0
		for elem := range channel3 {
			fmt.Println("Stage 3 received", elem)
			sum += elem
		}
		fmt.Println("Final Sum:", sum)
	}

	wg.Add(4)

	go initial(channel1, &wg)
	go stageOne(channel1, channel2, &wg)
	go stageTwo(channel2, channel3, &wg)
	go stageThree(channel3, &wg)

	wg.Wait()

}
