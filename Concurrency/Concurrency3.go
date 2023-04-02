package Concurrency

import (
	"fmt"
	"sync"
)

func ExerciseThree() {
	var counter int
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			mutex.Lock()
			counter++
			mutex.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", counter)
}
