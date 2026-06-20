package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(largestAltitude([]int{-5, 1, 5, 0, -7})) //[0,-5,-4,1,1,-6]
}

func largestAltitude(gain []int) int {

	var altitude []int

	altitude = append(altitude, 0)

	for _, altChange := range gain {
		point := altitude[len(altitude)-1] + altChange
		altitude = append(altitude, point)
	}

	fmt.Println(altitude)

	sort.Ints(altitude)

	return altitude[len(altitude)-1]
}

/*
func largestAltitude(gain []int) int {
	current, highest := 0, 0
	for _, g := range gain {
		current += g
		if current > highest {
			highest = current
		}
	}
	return highest
}
*/
