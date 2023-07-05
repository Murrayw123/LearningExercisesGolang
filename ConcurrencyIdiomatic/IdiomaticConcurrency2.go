package ConcurrencyIdiomatic

import (
	"fmt"
	"sync"
	"time"
)

func producer(c chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		c <- i // send integer to channel
		time.Sleep(time.Second)
	}
}

func consumer(id int, c <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range c { // receive from channel
		fmt.Printf("Consumer %d received %d\n", id, num)
	}
}

func ConcurrencyIdiomatic2() {
	producerWg := &sync.WaitGroup{}
	consumerWg := &sync.WaitGroup{}

	c := make(chan int)

	for i := 0; i < 10; i++ {
		producerWg.Add(1)
		go producer(c, producerWg) // We have multiple producers
	}

	consumerCount := 15
	consumerWg.Add(consumerCount)

	for i := 0; i < consumerCount; i++ {
		go consumer(i, c, consumerWg)
	}

	producerWg.Wait() // wait for all producers to finish
	close(c)          // it is safe to close the channel now
	consumerWg.Wait() // wait for all consumers to finish
}
