package main

import (
	"fmt"
)

func isValid(s string) bool {
	database := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	stack := []rune{}

	if len(s) <= 1 {
		return false
	}

	for _, char := range s {
		if _, exists := database[char]; exists {
			stack = append(stack, char)
		} else if len(stack) > 0 && database[stack[len(stack)-1]] == char {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}

func main() {
	result := isValid("()[]{}")
	fmt.Println(result) // Output: true

	result = isValid("(]")
	fmt.Println(result) // Output: false

	result = isValid("([)]")
	fmt.Println(result) // Output: false

	result = isValid("{[]}")
	fmt.Println(result) // Output: true
}
