package main

func main() {
	c := "abab"
	characterReplacement(c, 2)
}

func characterReplacement(s string, k int) int {

	if len(s) == 0 {
		return 0
	}

	maxCount, maxLength, charCount, start := 0, 0, make(map[byte]int), 0

	for end := 0; end < len(s); end++ {
		char := s[end]
		charCount[char]++
		maxCount = max(maxCount, charCount[char])

		if end-start+1-maxCount > k {
			charCount[s[start]]--
			start++
		}

		maxLength = max(maxLength, end-start+1)

	}
	return maxLength

}
