package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	// Hash Map for storing the character frequency
	requiredCharacters := make(map[byte]int)

	// Store the frequency of all the characters of string t
	for i := 0; i < len(t); i++ {
		requiredCharacters[t[i]]++
	}

	left, right := 0, 0
	minLength := math.MaxInt32
	minWindow := ""

	// Remaining characters to be found in the current window
	remaining := len(t)

	for right < len(s) {
		curChar := s[right]
		if requiredCharacters[curChar] > 0 {
			remaining--
		}

		requiredCharacters[curChar]--

		// Check if all characters are found in the current window
		for remaining == 0 {
			// Update the minimum window
			if right-left+1 < minLength {
				minLength = right - left + 1
				minWindow = s[left : right+1]
			}

			// Move the left pointer to find a smaller window
			leftChar := s[left]
			requiredCharacters[leftChar]++
			if requiredCharacters[leftChar] > 0 {
				remaining++
			}

			left++
		}

		// Move the right pointer to expand the window
		right++
	}

	return minWindow
}

func main() {
	searchString := "ADOBECODEBANC"
	t := "ABC"
	result := minWindow(searchString, t)
	fmt.Println(result)
}
