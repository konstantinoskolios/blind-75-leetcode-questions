// Space Complexity O(1)
// Time Complexity O(2n) => O(n)

package main

import "fmt"

func main() {
    x := []int{1, 9, 8, 4, 0, 0, 2, 7, 0, 6, 0}
    fmt.Printf("%v", pushZerosToEnd(x))
}

func pushZerosToEnd(nums []int) []int {
    l := 0
    for i, value := range nums {
		fmt.Printf("%v %v\n",i,l)
        if value != 0 {
            if i != l {
                nums[l] = nums[i]
				nums[i] = 0
            }
            l++
        }
    }
    return nums
}
