package main

import "fmt"

func main() {
	fmt.Println(combinationSum4([]int{1, 2, 3}, 4))
	fmt.Println(combinationSum4([]int{9}, 3))
}

func combinationSum4(nums []int, target int) int {
	dp := make(map[int]int)
	dp[0] = 1

	for total := 1; total <= target; total++ {
		dp[total] = 0
		for _, n := range nums {
			dp[total] += dp[total-n]
		}
	}


	return dp[target]
}