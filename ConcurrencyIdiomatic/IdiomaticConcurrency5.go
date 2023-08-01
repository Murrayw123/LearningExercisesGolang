package ConcurrencyIdiomatic

import (
	"fmt"
	"sync"
	"time"
)

func timeout(inputChannel chan string, timeoutDuration time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case msg, ok := <-inputChannel:
			if ok {
				fmt.Println(msg)
			} else {
				return
			}
		case <-time.After(timeoutDuration * time.Second):
			fmt.Println("Timeout occurred")
			return
		}
	}
}

func ConcurrencyIdiomaticFive() {
	inputChannel := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		time.Sleep(2 * time.Second)
		inputChannel <- "hey hey"
		time.Sleep(500 * time.Millisecond)
		inputChannel <- "Sending even more data baybeeee"
		fmt.Println("Closing")
		close(inputChannel)
	}()

	go timeout(inputChannel, 5, &wg)
	fmt.Println("Waiting")
	wg.Wait()
}
