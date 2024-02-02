package main

import (
	"fmt"
)

func isPalindrome(s string) bool {
	// This function checks whether a given string is a palindrome
	runeS := []rune(s)
	for i, j := 0, len(runeS)-1; i < j; i, j = i+1, j-1 {
		if runeS[i] != runeS[j] {
			return false
		}
	}
	return true
}

func palindromePairs(words []string) [][]int {
	// This function returns a list of indices that is a concatenation of two words and form a palindrome
	database := make(map[string][]int)

	for idx, word := range words {
		starter := idx + 1
		end := len(words)

		for jdx := starter; jdx < end; jdx++ {
			forComparison := words[jdx]
			concatenation := word + forComparison

			if isPalindrome(concatenation) && len(database[concatenation]) == 0 {
				database[concatenation] = []int{idx, jdx}
			} else if isPalindrome(concatenation) {
				database[concatenation] = append(database[concatenation], idx, jdx)
			}
		}
	}

	outerStarter := len(words) - 1
	for kdx := outerStarter; kdx >= 0; kdx-- {
		starter := kdx - 1
		end := 0

		for ldx := starter; ldx >= end; ldx-- {
			forComparison := words[ldx]
			concatenation := words[kdx] + forComparison

			if isPalindrome(concatenation) && len(database[concatenation]) == 0 {
				database[concatenation] = []int{kdx, ldx}
			} else if isPalindrome(concatenation) {
				database[concatenation] = append(database[concatenation], kdx, ldx)
			}
		}
	}

	var result [][]int

	for _, indices := range database {
		if len(indices) > 2 {
			for idx := 0; idx < len(indices); idx += 2 {
				result = append(result, []int{indices[idx], indices[idx+1]})
			}
		} else {
			result = append(result, indices)
		}
	}

	return result
}

func main() {
	words := []string{"bat", "tab", "cat"}
	fmt.Println(palindromePairs(words))

	words2 := []string{"abcd", "dcba", "lls", "s", "sssll"}
	fmt.Println(palindromePairs(words2))

	words3 := []string{"a", ""}
	fmt.Println(palindromePairs(words3))

	words4 := []string{"ci", "vic", "no", "on"}
	fmt.Println(palindromePairs(words4))

	words5 := []string{"baab", "abaabaa", "aaba", ""}
	fmt.Println(palindromePairs(words5))
}
