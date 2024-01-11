/*Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order
*/

package main

import "fmt"

func main() {
	input1 := []int{2,7,11,15}
	input2 := []int{3,2,4}
	input3 := []int{3,3}
	input4 := []int{2,2,5,11}
	target1 := 9
	target2 := 6
	target3 := 6
	target4 := 10
	result1 :=twoSum(input1, target1)
	result2 :=twoSum(input2, target2)
	result3 :=twoSum(input3, target3)
	result4 :=twoSum(input4, target4)
	printResult(result1)
	printResult(result2)
	printResult(result3)
	printResult(result4)
}

func printResult(result []int){
	fmt.Printf("Result: %v\n",result)
}

func twoSum(nums []int, target int) []int {

	sizeOf := len(nums)
	
	for i:=0; i< sizeOf; i++{
		for j:=i+1; j< sizeOf; j++{
			if(nums[i] + nums[j] == target){
				return []int{i,j}
			}
		}
	}

	return []int{}

	}
		

