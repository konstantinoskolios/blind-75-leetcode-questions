package main

import "fmt"

func main() {
	res := []int{-2, 3, -4}
	fmt.Printf("%v\n", maxProduct(res))
}

func maxProduct(nums []int) int {
    size := len(nums)

    if size == 0 {
        return 0
    }
    if size == 1 {
        return nums[0]
    }

    maxProduct := nums[0]
    minProduct := nums[0]
    result := nums[0]

    for i := 1; i < size; i++ {
        if nums[i] < 0 {
            maxProduct, minProduct = minProduct, maxProduct
        }

        maxProduct = max(nums[i], maxProduct*nums[i])
        minProduct = min(nums[i], minProduct*nums[i])

        result = max(result, maxProduct)
    }

    return result
}


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}