package ConcurrencyIdiomatic

import (
	"fmt"
	"sync"
	"time"
)

func listen(outputChannel chan int) {
	for i := range outputChannel {
		fmt.Println("output channel received", i)
	}
}

func fanOut(inputChannel chan int, outputChannels []chan int, wg *sync.WaitGroup) {
	defer func() {
		// input channel done
		wg.Done()
		// close each output channel and mark the waitgroup as done
		for _, ch := range outputChannels {
			close(ch)
			wg.Done()
		}
	}()

	sendTo := 0
	for i := range inputChannel {
		for v, ch := range outputChannels {
			if v == sendTo {
				fmt.Printf("sending channel %d value %d\n", v, i)
				ch <- i
			}
		}
		if sendTo < len(outputChannels) {
			sendTo++
		} else {
			sendTo = 0
		}
	}
}

func sendToInputChannel(inputChannel chan int) {
	defer close(inputChannel)
	for i := 0; i < 10; i++ {
		inputChannel <- i
		time.Sleep(500 * time.Millisecond)
	}
}

func ConcurrencyIdiomatic4() {
	inputChannel := make(chan int)
	outputChannels := make([]chan int, 3)

	for i := range outputChannels {
		outputChannels[i] = make(chan int)
	}

	wg := sync.WaitGroup{}
	wg.Add(len(outputChannels))

	for _, ch := range outputChannels {
		go listen(ch)
	}

	go fanOut(inputChannel, outputChannels, &wg)

	wg.Add(1)
	go sendToInputChannel(inputChannel)

	wg.Wait()
}
