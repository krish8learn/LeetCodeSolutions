package main

import "fmt"

func main() {

	fmt.Println(findDifference([]int{-68, -80, -19, -94, 82, 21, -43}, []int{-63}))
}

func findDifference(nums1 []int, nums2 []int) [][]int {

	num1Map, num2Map := make(map[int]bool, len(nums1)), make(map[int]bool, len(nums2))

	for _, value := range nums1 {
		num1Map[value] = true
	}

	for _, value := range nums2 {
		num2Map[value] = true
	}

	var only1, only2 []int
	for key := range num1Map {
		if !num2Map[key] {
			only1 = append(only1, key)
		}
	}

	for key := range num2Map {
		if !num1Map[key] {
			only2 = append(only2, key)
		}
	}

	return [][]int{only1, only2}
}
