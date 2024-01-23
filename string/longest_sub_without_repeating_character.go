package main

import "fmt"

func main() {
    fmt.Println(lengthOfLongestSubstring("abcabcbb")) 
    fmt.Println(lengthOfLongestSubstring("bbbbb"))    
    fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

func lengthOfLongestSubstring(s string) int {
    charSet := make(map[byte]bool)
    l := 0
    res := 0

    for r := 0; r < len(s); r++ {
        for charSet[s[r]] {
            delete(charSet, s[l])
            l++
        }
        charSet[s[r]] = true
        res = max(res, r-l+1)
    }

    return res
}
