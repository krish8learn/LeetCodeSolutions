package main

func main() {

	res := findTheDifference("abcd", "abcde")
	println(string(res))
}

func findTheDifference(s string, t string) byte {
	var x byte = 0
	for i := 0; i < len(s); i++ {
		x ^= s[i]
	}
	for i := 0; i < len(t); i++ {
		x ^= t[i]
	}
	return x
}
