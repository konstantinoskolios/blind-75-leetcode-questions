package main

import "fmt"

func main() {
    fmt.Println(lengthOfLongestSubstring("abcabcbb")) 
    fmt.Println(lengthOfLongestSubstring("bbbbb"))    
    fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

func lengthOfLongestSubstring(s string) int {
    charIndex := make(map[byte]int)
    maxLength, start := 0, 0

    for end := 0; end < len(s); end++ {
        char := s[end]
        if index, found := charIndex[char]; found && index >= start {
            start = index + 1
        }

        charIndex[char] = end
        maxLength = max(maxLength, end-start+1)
    }

    return maxLength
}
