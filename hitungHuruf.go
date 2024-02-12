package main

import (
	"fmt"
	"unicode"
)

func hitungHuruf(s string) map[rune]int {
	count := make(map[rune]int)

	for _, char := range s {
		if 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' {
			char = unicode.ToLower(char)
			count[char]++
		}
	}
	return count
}

func main() {
	words := "coding is fun"
	wordsCount := hitungHuruf(words)

	for letter, count := range wordsCount {
		fmt.Printf("%c: %d\n", letter, count)
	}
}
