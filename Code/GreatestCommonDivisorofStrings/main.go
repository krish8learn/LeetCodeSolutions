package main

import "fmt"

func main() {

	fmt.Println(gcdOfStrings("ABCABC", "ABC"))
}

func gcdOfStrings(str1 string, str2 string) string {

	if str1+str2 != str2+str1 {
		// no common divisor
		return ""
	}

	divisor := gcd(len(str1), len(str2))

	return str1[:divisor]
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}
