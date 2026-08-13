package main

import (
	"fmt"
	"time"
)

func main() {
	nasdaq := make(chan string)
	nyse := make(chan string)

	// nasdaq pushing data
	go func() {
		defer close(nasdaq)
		nasdaq <- "Nasdaq data 110"
		time.Sleep(10 * time.Millisecond)
		nasdaq <- "Nasdaq data 111"
	}()

	// nyse data
	go func() {
		defer close(nyse)
		nyse <- "Nyse Data 100"
		time.Sleep(10 * time.Millisecond)
		nyse <- "Nyse Data 101"
	}()

	feed := MergeFeed(nasdaq, nyse)

	for tick := range feed {
		fmt.Println(tick)
	}
}

func MergeFeed(nasdaq, nyse <-chan string) <-chan string {

	output := make(chan string)

	go func() {
		defer close(output)

		for nasdaq != nil || nyse != nil {
			select {
			case tick, ok := <-nasdaq:
				if !ok {
					// channel is closed,set to nil
					// in select statement reading it from nil channel blocks forever
					nasdaq = nil
					continue
				}
				output <- "NASDAQ: " + tick
			case tick, ok := <-nyse:
				if !ok {
					nyse = nil
					continue
				}
				output <- "NYSE: " + tick
			}
		}
	}()

	return output
}
