package main

import "fmt"

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
}

func pivotIndex(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	leftSum := 0
	for i, n := range nums {
		if leftSum == total-leftSum-n {
			return i
		}
		leftSum += n
	}
	return -1
}
