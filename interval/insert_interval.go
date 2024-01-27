package main

import (
	"fmt"
)

func main() {
	// res := insert([][]int{{6,9},{1,3}},[]int{2,5})
	res := insert([][]int{{1,2},{3,5},{6,7},{8,10},{12,16}},[]int{4,8})
	fmt.Println(res)
}

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(newInterval) == 0 {
		return intervals
	}

	start, end := newInterval[0], newInterval[1]
	result := make([][]int, 0)

	inserted := false

	for _, v := range intervals {
		start_b, end_b := v[0], v[1]

		if end < start_b && !inserted {
			result = append(result, []int{start, end})
			inserted = true
		}

		if (start <= end_b) && (end >= start_b) {
			start = min(start, start_b)
			end = max(end, end_b)
		} else {
			result = append(result, []int{start_b, end_b})
		}
	}

	if !inserted {
		result = append(result, []int{start, end})
	}

	return result
}