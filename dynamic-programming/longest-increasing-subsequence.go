package main

import "fmt"

func main() {
    fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
    fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
    fmt.Println(lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))
}

func lengthOfLIS(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    dp := make([]int, 0)

    for _, num := range nums {
        left, right := 0, len(dp)-1
        for left <= right {
            mid := left + (right-left)/2
            if dp[mid] < num {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }

        if left == len(dp) {
            dp = append(dp, num)
        } else {
            dp[left] = num
        }
    }

    return len(dp)
}
