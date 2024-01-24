package main

import (
	"fmt"
)

func minWindow(s string, t string) string {
	const MaxInt = int(^uint(0) >> 1)
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	targetMap := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		targetMap[t[i]]++
	}

	left, right := 0, 0
	minWindowStart, minWindowEnd := 0, MaxInt
	requiredChars := len(targetMap)

	for right < len(s) {
		if count, exists := targetMap[s[right]]; exists {
			if count > 0 {
				requiredChars--
			}
			targetMap[s[right]]--
		}

		for requiredChars == 0 {
			if right-left < minWindowEnd-minWindowStart {
				minWindowStart, minWindowEnd = left, right
			}

			if count, exists := targetMap[s[left]]; exists {
				targetMap[s[left]]++
				if count == 0 {
					requiredChars++
				}
			}

			left++
		}

		right++
	}

	if minWindowEnd == MaxInt {
		return ""
	}

	return s[minWindowStart : minWindowEnd+1]
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	result := minWindow(s, t)
	fmt.Println("Output:", result)
}
