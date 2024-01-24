package main

import (
	"fmt"
)

func main() {
	var input1 = "abc"
	var input2 = "aaa"
	fmt.Printf("%v\n",countSubstrings(input1))
	fmt.Printf("%v\n",countSubstrings(input2))
}

func countSubstrings(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		count += expandAroundCenter(s, i, i) //odd
		count += expandAroundCenter(s, i, i+1) //even
	}
	return count
}

func expandAroundCenter(s string, left, right int) int {
	count := 0
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
		count++
	}
	return count
}