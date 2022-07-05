package leetcode

import "strings"

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	for i, word := range words {
		words[i] = reverseWord(word)
	}
	result := strings.Join(words, " ")
	return result
}

func reverseWord(s string) (reversed string) {
	for _, char := range s {
		reversed = string(char) + reversed
	}
	return // returns reversed (go)
}
