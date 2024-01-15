package main

import (
	"fmt"
	"sort"
)

var input1 = []int{1, 2, 3, 1}
var input2 = []int{1, 2, 3, 4}
var input3 = []int{1, 1, 1, 3, 3, 4, 3, 4, 2}
var input4 = []int{3, 3}
var input5 = []int{3}
var input6 = []int{3,1}

func main() {
	fmt.Println(containsDuplicate(input1))
	fmt.Println(containsDuplicate(input2))
	fmt.Println(containsDuplicate(input3))
	fmt.Println(containsDuplicate(input4))
	fmt.Println(containsDuplicate(input5))
	fmt.Println(containsDuplicate(input6))
}

func containsDuplicate(nums []int) bool {
    freqMap := make(map[int]int,len(nums))

    for _, num := range nums {
        if freqMap[num] > 0 {
            return true
        }
        
        freqMap[num]++
    }

    return false
}