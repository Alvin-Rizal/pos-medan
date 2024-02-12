package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func countLetters(words []string) map[rune]int {
	count := make(map[rune]int)
	for _, word := range words {
		for _, char := range word {
			count[char]++
		}
	}
	return count
}

func urutkanHuruf(words []string) string {
	count := countLetters(words)
	uniqueLetters := make([]rune, 0, len(count))
	for letter := range count {
		uniqueLetters = append(uniqueLetters, letter)
	}

	sort.Slice(uniqueLetters, func(i, j int) bool {
		freqI := count[uniqueLetters[i]]
		freqJ := count[uniqueLetters[j]]

		if freqI != freqJ {
			return freqI > freqJ
		}

		runeI := uniqueLetters[i]
		runeJ := uniqueLetters[j]

		isUpperI := unicode.IsUpper(runeI)
		isUpperJ := unicode.IsUpper(runeJ)

		if isUpperI == isUpperJ {
			return runeI < runeJ
		}
		return isUpperI
	})

	var sorted strings.Builder
	for _, letter := range uniqueLetters {
		sorted.WriteRune(letter)
	}
	return sorted.String()
}

func main() {
	words := []string{"Abc", "bCd"}
	words2 := []string{"Oke", "One"}
	words3 := []string{"Pendanaan", "Terproteksi", "Untuk", "Dampak", "Berarti"}

	sortedLetters := urutkanHuruf(words)
	sortedLetters2 := urutkanHuruf(words2)
	sortedLetters3 := urutkanHuruf(words3)

	fmt.Println("Hasilnya : ", sortedLetters)
	fmt.Println("Hasilnya : ", sortedLetters2)
	fmt.Println("Hasilnya : ", sortedLetters3)
}
