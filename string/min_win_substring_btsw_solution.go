package main

import (
	"fmt"
	"math"
)

type Solution struct{}

func (sol *Solution) minWindow(searchString string, t string) string {
	n := len(searchString)

	// It contains the min length seen so far
	minWindowSize := math.MaxInt32

	// It contains the minimum length substring
	var minWindow string

	// The nested for loop helps us generate all the substrings
	for left := 0; left < n; left++ {
		for right := left; right < n; right++ {

			// Generate the substring
			windowSnippet := searchString[left : right+1]

			// Check if the generated char contains all the characters of target
			windowContainsAll := sol.containsAll(windowSnippet, t)

			// If it is a valid substring
			// And the length is less than the minimum so far
			// Update the answer
			if windowContainsAll && len(windowSnippet) < minWindowSize {
				minWindowSize = len(windowSnippet)
				minWindow = windowSnippet
			}
		}
	}

	return minWindow
}

// Helper function to check if all the char of string are
// Present in the string searchString
func (sol *Solution) containsAll(searchString string, t string) bool {
	requiredCharacters := make(map[byte]int)

	for i := 0; i < len(t); i++ {
		occurrences := 0
		if val, exists := requiredCharacters[t[i]]; exists {
			occurrences = val
		}

		requiredCharacters[t[i]] = occurrences + 1
	}

	for i := 0; i < len(searchString); i++ {
		curr := searchString[i]

		if occurrences, exists := requiredCharacters[curr]; exists {
			// Calculate what the new occurrence count will be
			newOccurrences := occurrences - 1

			// If we have satisfied all of the characters for this character, remove the key
			// from the hashtable.
			// Otherwise, just update the mapping with 1 less occurrence to satisfy for
			if newOccurrences == 0 {
				delete(requiredCharacters, curr)
			} else {
				requiredCharacters[curr] = newOccurrences
			}
		}
	}

	// If we satisfied all characters, the required characters hashtable will be empty
	return len(requiredCharacters) == 0
}

func main() {
	sol := &Solution{}
	searchString := "ADOBECODEBANC"
	t := "ABC"
	result := sol.minWindow(searchString, t)
	fmt.Println(result)
}
