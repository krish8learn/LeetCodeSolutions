package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// this will simulate the flaky API
func FetchPrice(requestId string) float64 {

	fmt.Println("received rq", requestId)
	randomLatency := time.Duration(rand.Intn(81)+20) * time.Millisecond
	time.Sleep(randomLatency)

	return 133.567
}

func GetPriceWithTimeOut(rqid string) (float64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	// main routine and extra go routine must communicate through a channel, which must be buffered, for unlimited channel this gets block
	reschn := make(chan float64, 1)

	go func() {
		price := FetchPrice(rqid)
		reschn <- price
	}()

	// use select to catch possible condition
	select {
	case price := <-reschn:
		return price, nil
	case <-ctx.Done():
		return 0, fmt.Errorf("timeout")
	}
}

func main() {
	for index := range 10 {
		price, err := GetPriceWithTimeOut(fmt.Sprintf("req%d", index))
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(price)
		}
	}
}
