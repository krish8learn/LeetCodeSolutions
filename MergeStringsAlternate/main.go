package main

import "fmt"

func main() {
	fmt.Println(mergeAlternately("abc", "pqr")) // apbqcr
}

func mergeAlternately(word1, word2 string) string {

	result, index := "", 0
	len1, len2 := len(word1), len(word2)

	for index < len1 || index < len2 {
		if index < len1 {
			result += string(word1[index])
		}

		if index < len2 {
			result += string(word2[index])
		}

		index++
	}

	return result
}
