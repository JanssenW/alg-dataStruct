package main

import "fmt"

func main() {

	s := "abcabcbb"
	maxLen := maximumLengthSubstring(s)
	fmt.Printf("Max length substring is: %d\n", maxLen)
}

func maximumLengthSubstring(s string) int {
	r, l := 0, 0
	maxLen := 0
	sRune := []rune(s)
	m := make(map[rune]int)

	for r < len(sRune) {
		m[sRune[r]]++

		for m[sRune[r]] > 2 {
			m[sRune[l]]--
			l++
		}

		maxLen = max(maxLen, r-l+1)
		r++
	}

	return maxLen
}
