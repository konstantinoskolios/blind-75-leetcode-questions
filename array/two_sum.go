package main

import "fmt"

func main() {
	input := []int{3,2, 4}
	target := 6
	result := twoSum(input, target)
	printResult(result)
}

func printResult(result []int) {
	fmt.Printf("Result: %v\n", result)
}

//O(n)
func twoSum(nums []int, target int) []int {
	ht := make(map[int]int, len(nums))

	for i, v := range nums {
		complement := target - v
		if index, ok := ht[complement]; ok {
			return []int{index, i}
		}
		ht[v] = i
	}

	return []int{}
}

//O(n^2)
// func twoSum(nums []int, target int) []int {

// 	sizeOf := len(nums)
	
// 	for i:=0; i< sizeOf; i++{
// 		for j:=i+1; j< sizeOf; j++{
// 			if(nums[i] + nums[j] == target){
// 				return []int{i,j}
// 			}
// 		}
// 	}

// 	return []int{}

// 	}
