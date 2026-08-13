package main

import "strings"

func main() {
	s := "abab"
	println(repeatedSubstringPattern(s))
}

// Uses the classic trick: s is a repeated substring iff s appears
// in (s+s)[1:len(s+s)-1]. This is concise and interview-friendly.
func repeatedSubstringPattern(s string) bool {
	if len(s) == 0 {
		return false
	}
	ss := s + s
	return strings.Contains(ss[1:len(ss)-1], s)
}
