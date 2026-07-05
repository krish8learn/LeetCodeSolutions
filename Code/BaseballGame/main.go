package main

import "strconv"

func main() {
	operations := []string{"5", "-2", "4", "C", "D", "9", "+", "+"}
	result := calPoints(operations)
	println(result) // Output: 30 it should be 27

}

func calPoints(operations []string) int {

	var stack []int

	for _, op := range operations {
		switch op {
		case "+":
			stack = append(stack, stack[len(stack)-1]+stack[len(stack)-2])
		case "D":
			stack = append(stack, 2*stack[len(stack)-1])
		case "C":
			stack = stack[:len(stack)-1]
		default:
			num, _ := strconv.Atoi(op)
			stack = append(stack, num)
		}
	}

	sum := 0
	for _, score := range stack {
		sum += score
	}
	return sum
}
