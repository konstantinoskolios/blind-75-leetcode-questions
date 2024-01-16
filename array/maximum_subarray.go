package main

import "fmt"

func main() {
	array := []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Printf("res: %v\n",maxSubArray(array))
}


func maxSubArray(nums []int) int {
    size := len(nums)

    if size == 0 {
        return 0
    }
    if size == 1 {
        return nums[0]
    }

    maxSum := nums[0]
    currentSum := nums[0]

    for i := 1; i < size; i++ {
        currentSum = max(nums[i], currentSum+nums[i])
        maxSum = max(maxSum, currentSum)
    }

    return maxSum
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
