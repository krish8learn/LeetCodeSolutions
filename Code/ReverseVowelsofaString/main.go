package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	fmt.Println(reverseVowels("IceCreAm"))
}

const vowelString = "aeiou"

func reverseVowels(s string) string {
	chars := []rune(s)
	vowels := make([]rune, 0)
	positions := make([]int, 0)

	for i, ch := range chars {
		if strings.ContainsRune(vowelString, unicode.ToLower(ch)) {
			vowels = append(vowels, ch)
			positions = append(positions, i)
		}
	}

	for i, pos := range positions {
		chars[pos] = vowels[len(vowels)-1-i]
	}

	return string(chars)
}
