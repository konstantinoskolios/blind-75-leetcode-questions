package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 1}
	fmt.Println(rob(nums1))

	nums2 := []int{2, 7, 9, 3, 1}
	fmt.Println(rob(nums2))
}


func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	}

	robCurrent, skipCurrent := nums[0], 0

	for i := 1; i < n; i++ {
	
		newRobCurrent := skipCurrent + nums[i]
		newSkipCurrent := max(robCurrent, skipCurrent)
		robCurrent, skipCurrent = newRobCurrent, newSkipCurrent
	}

	return max(robCurrent, skipCurrent)
}

