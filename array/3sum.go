package main

import (
	"fmt"
	"sort"
)

func main() {
	res := []int{-1, 0, 1, 2, -1, -4}
	fmt.Printf("%v\n", threeSum(res))

}

func threeSum(nums []int) [][]int {

	result := [][]int{}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := 0; i < len(nums)-1; i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target := -nums[i]
		l, r := i+1, len(nums)-1

		for l < r {
			sum := nums[l] + nums[r]

			if sum == target {
				result = append(result, []int{nums[i], nums[r], nums[l]})
				l++
				i--

				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else if sum < target {
				l++
			} else {
				r--
			}
		}
	}

	return result
}
