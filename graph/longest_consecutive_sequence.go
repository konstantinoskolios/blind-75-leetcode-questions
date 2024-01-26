package main

import (
	"fmt"
	"sort"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)

	maxSequence := 1
	currentSequence := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			if nums[i] == nums[i-1]+1 {
				currentSequence++
			} else if nums[i] != nums[i-1] {
				currentSequence = 1
			}

			maxSequence = max(maxSequence, currentSequence)
		}
	}

	return maxSequence
}


func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))      // Output: 4
	fmt.Println(longestConsecutive([]int{ 3, 7, 2, 5, 8, 4, 6, 1,0,0})) // Output: 9
	fmt.Println(longestConsecutive([]int{ 0,-1})) // Output: 2
}
