package main

import (
	"fmt"
)

func main() {
	var input1 = []int{1, 2, 3, 4}
	var input2 = []int{-1, 1, 0, -3, 3}
	var input3 = []int{0, 0}
	var input4 = []int{4, 3, 2, 1, 2}
	var input5 = []int{5,9,2,-9,-9,-7,-8,7,-9,10}
	fmt.Printf("%v\n", productExceptSelf(input1))
	fmt.Printf("%v\n", productExceptSelf(input2))
	fmt.Printf("%v\n", productExceptSelf(input3))
	fmt.Printf("%v\n", productExceptSelf(input4))
	fmt.Printf("%v\n", productExceptSelf(input5))
}


func productExceptSelf(nums []int) []int {
    n := len(nums)
    result := make([]int, n)

    leftProduct := 1
    for i := 0; i < n; i++ {
        result[i] = leftProduct
        leftProduct *= nums[i]
    }

	rightProduct := 1
    for i := n - 1; i >= 0; i-- {
        result[i] *= rightProduct
        rightProduct *= nums[i]
    }

    return result
}
