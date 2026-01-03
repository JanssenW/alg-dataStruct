package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "Abc Def Ghi j"
	fmt.Println(reverseWords(s))
}

func reverseWords(s string) string {

	words := strings.Split(s, " ")

	for k, word := range words {

		words[k] = reverseWord(word)

	}

	return strings.Join(words, " ")

}

func reverseWord(s string) string {
	wordRune := []rune(s)

	//Two Pointers in Arrays
	for i, j := 0, len(wordRune)-1; i < j; i, j = i+1, j-1 {
		wordRune[i], wordRune[j] = wordRune[j], wordRune[i]
	}

	return string(wordRune)
}
