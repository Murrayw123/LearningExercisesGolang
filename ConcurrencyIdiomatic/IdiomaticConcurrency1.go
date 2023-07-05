package ConcurrencyIdiomatic

import (
	"fmt"
	"sync"
)

func fanIn(c1, c2 <-chan int) <-chan int {
	var wg sync.WaitGroup

	out := make(chan int)

	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			out <- n
		}
	}

	wg.Add(2)
	go output(c1)
	go output(c2)

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func ConcurrencyIdiomatic1() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c1 <- i
		}
		close(c1)
	}()

	go func() {
		for i := 10; i < 20; i++ {
			c2 <- i
		}
		close(c2)
	}()

	for n := range fanIn(c1, c2) {
		fmt.Println(n)
	}
}
