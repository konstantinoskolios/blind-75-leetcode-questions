package main

import "fmt"

func main() {
	s1 := "leetcode"
	wordDict1 := []string{"leet", "code"}
	fmt.Println(wordBreak(s1, wordDict1))

	s2 := "applepenapple"
	wordDict2 := []string{"apple", "pen"}
	fmt.Println(wordBreak(s2, wordDict2))

	s3 := "catsandog"
	wordDict3 := []string{"cats", "dog", "sand", "and", "cat"}
	fmt.Println(wordBreak(s3, wordDict3)) 
}

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && contains(wordDict, s[j:i]) {
				dp[i] = true
				break
			}
		}
	}

	return dp[n]
}

func contains(wordDict []string, word string) bool {
	for _, w := range wordDict {
		if w == word {
			return true
		}
	}
	return false
}
