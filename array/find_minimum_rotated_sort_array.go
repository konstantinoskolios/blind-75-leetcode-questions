package main

import (
	"fmt"
)

var test1 = []int{3,4,5,1,2}
var test2 = []int{4,5,6,7,0,1,2}
var test3 = []int{11,13,15,17}

func main() {
	fmt.Printf("%v\n",findMin(test1))
	fmt.Printf("%v\n",findMin(test2))
	fmt.Printf("%v\n",findMin(test3))
}

func findMin(nums []int) int {
	l, r := 0, len(nums) -1

	for l < r {
		m := (l+r) /2
		if nums[m] > nums[r] {
			l = m + 1
		}else{
			r = m
		}
	}

	return nums[l]
}