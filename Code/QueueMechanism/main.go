package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	var lock sync.RWMutex
	// data push
	Publisher("krish", "name", &lock)
	Publisher("123", "number", &lock)

	// data consume subscribed topic only

	wg.Add(3)
	go Consumer("name", &wg, &lock)
	go Consumer("number", &wg, &lock)
	go Consumer("name", &wg, &lock)

	wg.Wait()
}

// queue
var List = make(map[string][]string)

func Publisher(data string, topic string, lock *sync.RWMutex) error {
	// pushing the data into queue
	lock.Lock()
	List[topic] = append(List[topic], data)
	lock.Unlock()
	return nil
}

func Consumer(topic string, wg *sync.WaitGroup, lock *sync.RWMutex) error {

	lock.Lock()
	queue := List[topic]
	for index, value := range queue {
		fmt.Println("data value and index", value, index)
		delete(List, topic)
	}
	lock.Unlock()

	wg.Done()
	return nil
}
