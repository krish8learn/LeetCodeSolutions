package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// using a buffered channel to prvent blocking of sender
	tickStream := make(chan string, 100)

	// simulate upstream web socket sending the ticket
	go func() {

		for index := range 20 {
			tickStream <- fmt.Sprintf("price is 150.%02d", index)
		}

		// close the channel and workers to terminate once channel is empty
		close(tickStream)
	}()

	// start processing with worker of 5
	StartTickProcessor(tickStream, 5)
}

func StartTickProcessor(tickStream <-chan string, numworkers int) {

	var wg sync.WaitGroup
	// spawn the specific number of go routines
	for index := range numworkers {
		wg.Add(1)
		go func(workerId int) {
			defer wg.Done()
			for tick := range tickStream {

				fmt.Println("worker processing to write in ", workerId, tick)
				time.Sleep(10 * time.Millisecond)
			}
		}(index)
	}

	wg.Wait()
	fmt.Println("All workers are done")
}
