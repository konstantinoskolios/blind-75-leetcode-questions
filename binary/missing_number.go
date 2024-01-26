package main

import (
	"fmt"
)

func main() {
	x := missingNumber([]int{3,0,2})
	fmt.Printf("%v\n",x)
}

func missingNumber(nums []int) int {
    sum1 := 0
    for _, v := range nums {
        sum1 += v
    }

    n := len(nums)
    sum2 := n * (n + 1) / 2

    return sum2 - sum1
}
