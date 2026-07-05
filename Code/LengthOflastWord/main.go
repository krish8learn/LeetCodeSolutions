package main

import "strings"

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}
	words := strings.Split(s, " ")
	return len(words[len(words)-1])
}
