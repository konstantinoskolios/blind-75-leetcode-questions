package main

import "fmt"

func main() {
	fmt.Println(climbStairs(3))
}

func climbStairs(n int) int {
	one, two := 1, 1
	for i := 0; i < n-1; i++ {
		temp := one
		one = one + two
		two = temp
	}
	return one
}
