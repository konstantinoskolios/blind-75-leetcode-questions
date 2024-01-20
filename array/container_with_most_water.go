package main

import (
	"fmt"
)

func main() {

	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	maxArea(input)
}

func maxArea(height []int) int {


	l, r, res := 0, len(height)-1, 0

	for l < r {
		area := (r - l) * min(height[l], height[r])
		res = max(res, area)

		if height[l] < height[r] {
			l++
		} else {
			r--
		}

	}

	fmt.Printf("%v\n", res)
	return res

}
