package main

import (
	"fmt"
	"sort"
)

func main() {
	input := [][]int{{1, 3}, {1, 10}, {2, 6}, {15, 18}}
	result := merge(input)
	fmt.Printf("%v\n", result)
}

func merge(input [][]int) [][]int {
	if len(input) == 0 {
		return [][]int{}
	}

	sort.Slice(input, func(i, j int) bool {
		return input[i][0] < input[j][0]
	})

	arr := [][]int{input[0]}
	x := 0

	fmt.Printf("%v\n",input)

	for i := 1; i < len(input); i++ {
		endLeft := arr[x][1]
		

		startRight, endRight := input[i][0], input[i][1]
		fmt.Printf("%v - %v\n",endLeft,startRight)
		if endLeft >= startRight {
			arr[x][1] = max(endLeft, endRight)
		} else {
			arr = append(arr, input[i])
			x++
		}
	}

	return arr
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


